package account

import (
	"backend/api/http/requests/accountrequests"
	"backend/api/pkg/customerrors"
	"backend/api/pkg/models"
	"backend/api/pkg/user"
	"backend/api/pkg/utils"
	"context"
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
	) ([]*models.Account, error)
	Edit(
		ctx context.Context,
		userID primitive.ObjectID,
		account *accountrequests.Edit,
	) error
}

type service struct {
	repo        Repository
	userService user.Service
}

func NewService(
	repository Repository,
) Service {
	return &service{
		repo: repository,
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
) ([]*models.Account, error) {
	return s.repo.GetByUserID(ctx, userID)
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
