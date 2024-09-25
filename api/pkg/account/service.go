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
		return err
	}

	return s.repo.Create(
		ctx,
		userID,
		account,
	)
}

func (s *service) GetByUserID(
	ctx context.Context,
	userID primitive.ObjectID,
) ([]*dtos.Account, error) {
	accounts, err := s.repo.GetByUserID(ctx, userID)
	if err != nil {
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

	return accountDtos, nil
}

func (s *service) CreateMany(
	ctx context.Context,
	userID primitive.ObjectID,
	accounts []*accountrequests.Create,
) error {
	userObj, err := s.checkAccountLimits(ctx, userID)
	if err != nil {
		return err
	}
	accountsCount, err := s.repo.CountByUserID(ctx, userID)
	if err != nil {
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
		return err
	}

	if notCreated > 0 {
		return customerrors.NewPartialAccountsError(created, notCreated)
	}
	return nil
}

func (s *service) Edit(
	ctx context.Context,
	userID primitive.ObjectID,
	account *accountrequests.Edit,
) error {
	return s.repo.Edit(ctx, userID, account)
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
