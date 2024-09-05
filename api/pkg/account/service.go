package account

import (
	"backend/api/pkg/customerrors"
	"backend/api/pkg/user"
	"backend/api/pkg/utils"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service interface {
	Create(
		ctx context.Context,
		userID primitive.ObjectID,
		token string,
		proxyType string,
		proxyLogin string,
		proxyPass string,
		proxyIP string,
		proxyPort string,
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

func (s *service) Create(
	ctx context.Context,
	userID primitive.ObjectID,
	token string,
	proxyType string,
	proxyLogin string,
	proxyPass string,
	proxyIP string,
	proxyPort string,
) error {
	userObj, err := s.userService.GetByID(ctx, userID)

	if err != nil {
		return err
	}

	if userObj.SubEnd < utils.GetDateTime() {
		return customerrors.ErrSubNotActive
	}

	accountsCount, err := s.repo.CountByUserID(ctx, userID)
	if err != nil {
		return err
	}

	if accountsCount >= userObj.MaxAccounts {
		return customerrors.ErrAccountsLimit
	}

	return s.repo.Create(
		ctx,
		userID,
		token,
		proxyType,
		proxyLogin,
		proxyPass,
		proxyIP,
		proxyPort,
	)
}
