package account

import (
	"backend/api/http/requests/accountrequests"
	activationPkg "backend/api/pkg/activation"
	"backend/api/pkg/customerrors"
	"backend/api/pkg/dtos"
	"backend/api/pkg/models"
	promoCodePkg "backend/api/pkg/promocode"
	"backend/api/pkg/user"
	"backend/api/pkg/utils"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

type Service interface {
	Create(
		ctx context.Context,
		userID primitive.ObjectID,
		account *accountrequests.Create,
	) error
	CreateMany(
		ctx context.Context,
		userID primitive.ObjectID,
		accounts []*accountrequests.Create,
	) error
	GetByUserID(
		ctx context.Context,
		userID primitive.ObjectID,
	) ([]*dtos.Account, error)
	Edit(
		ctx context.Context,
		userID primitive.ObjectID,
		account *accountrequests.Edit,
	) error
}

type service struct {
	repo              Repository
	userService       user.Service
	promoCodeService  promoCodePkg.Service
	activationService activationPkg.Service
}

func NewService(
	repository Repository,
	userService user.Service,
	promoCodeService promoCodePkg.Service,
	activationService activationPkg.Service,
) Service {
	return &service{
		repo:              repository,
		userService:       userService,
		promoCodeService:  promoCodeService,
		activationService: activationService,
	}
}

// Проверка условий подписки и лимита аккаунтов
func (s *service) checkAccountLimits(
	ctx context.Context,
	userID primitive.ObjectID,
) (*models.User, error) {
	userObj, err := s.userService.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	if userObj.SubEnd < utils.GetDateTime() {
		return nil, customerrors.ErrSubNotActive
	}

	accountsCount, err := s.repo.CountByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	if accountsCount >= userObj.MaxAccounts {
		return nil, customerrors.ErrAccountsLimit
	}

	return userObj, nil
}

func (s *service) Create(
	ctx context.Context,
	userID primitive.ObjectID,
	account *accountrequests.Create,
) error {
	_, err := s.checkAccountLimits(ctx, userID)
	if err != nil {
		log.Printf("create account: limits check failed: %v", err)
		return err
	}

	if err := s.repo.Create(
		ctx,
		userID,
		account,
	); err != nil {
		log.Printf("create account: repository error: %v", err)
		return err
	}
	log.Printf("account created for user %s", userID.Hex())
	return nil
}

func (s *service) GetByUserID(
	ctx context.Context,
	userID primitive.ObjectID,
) ([]*dtos.Account, error) {
	accounts, err := s.repo.GetByUserID(ctx, userID)
	if err != nil {
		log.Printf("get accounts: repository error: %v", err)
		return nil, err
	}

	var accountDtos []*dtos.Account
	for _, account := range accounts {
		lastActivation, err := s.getActivationForAccount(ctx, account.ID)
		if err != nil {
			return nil, err
		}

		proxyString := utils.BuildAccountProxyString(account)

		accountDto := &dtos.Account{
			Account:        *account,
			Proxy:          proxyString,
			LastActivation: lastActivation,
		}
		accountDtos = append(accountDtos, accountDto)
	}

	log.Printf("retrieved %d accounts for user %s", len(accountDtos), userID.Hex())
	return accountDtos, nil
}

func (s *service) CreateMany(
	ctx context.Context,
	userID primitive.ObjectID,
	accounts []*accountrequests.Create,
) error {
	userObj, err := s.checkAccountLimits(ctx, userID)
	if err != nil {
		log.Printf("create many accounts: limits check failed: %v", err)
		return err
	}
	accountsCount, err := s.repo.CountByUserID(ctx, userID)
	if err != nil {
		log.Printf("create many accounts: count error: %v", err)
		return err
	}

	// Рассчитываем, сколько профилей можно создать
	availableSlots := userObj.MaxAccounts - accountsCount

	// Ограничиваем создание профилей по доступному количеству слотов
	accountsToCreate := accounts
	if len(accounts) > availableSlots {
		accountsToCreate = accounts[:availableSlots]
	}
	created := len(accountsToCreate)
	notCreated := len(accounts) - created

	// Создаем профили
	err = s.repo.CreateMany(
		ctx,
		userID,
		accountsToCreate,
	)
	if err != nil {
		log.Printf("create many accounts: repository error: %v", err)
		return err
	}

	if notCreated > 0 {
		log.Printf("create many accounts: %d accounts not created due to limits", notCreated)
		return customerrors.NewPartialAccountsError(created, notCreated)
	}
	log.Printf("created %d accounts for user %s", created, userID.Hex())
	return nil
}

func (s *service) Edit(
	ctx context.Context,
	userID primitive.ObjectID,
	account *accountrequests.Edit,
) error {
	if err := s.repo.Edit(ctx, userID, account); err != nil {
		log.Printf("edit account %s: %v", account.ID.Hex(), err)
		return err
	}
	log.Printf("account %s edited", account.ID.Hex())
	return nil
}

func (s *service) getActivationForAccount(
	ctx context.Context,
	accountID primitive.ObjectID,
) (*dtos.Activation, error) {
	lastActivation, err := s.activationService.GetLastByAccountID(ctx, accountID)
	if err != nil {
		if errors.Is(err, customerrors.ErrActivationNotFound) {
			return nil, nil
		}
		return nil, err
	}

	promoCode, err := s.promoCodeService.GetByID(ctx, lastActivation.PromoCodeID)
	if err != nil && !errors.Is(err, customerrors.ErrPromoCodeNotFound) {
		return nil, err
	}

	return &dtos.Activation{
		Activation: *lastActivation,
		PromoCode:  promoCode,
	}, nil
}
