package user

import (
	"backend/api/pkg/customerrors"
	"backend/api/pkg/models"
	"backend/api/pkg/utils"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service interface {
	Register(ctx context.Context, mnemonic []string) (string, error)
	Login(ctx context.Context, mnemonic []string) (string, error)
	GetByID(ctx context.Context, id primitive.ObjectID) (*models.User, error)
}

type service struct {
	repo Repository
}

func NewService(
	repository Repository,
) Service {
	return &service{
		repo: repository,
	}
}

func (s *service) Register(ctx context.Context, mnemonic []string) (string, error) {
	if utils.ValidateMnemonic(mnemonic) {
		return "", customerrors.ErrInvalidMnemonic
	}
	seed := utils.MnemonicToSeed(mnemonic)
	_, err := s.repo.GetBySeed(ctx, seed)
	if err == nil {
		return "", customerrors.ErrUserAlreadyExists
	} else if !errors.Is(err, customerrors.ErrUserNotFound) {
		return "", err
	}
	userID, err := s.repo.Register(ctx, seed)
	if err != nil {
		return "", err
	}

	return utils.EncodeJWT(userID)
}

func (s *service) Login(ctx context.Context, mnemonic []string) (string, error) {
	if utils.ValidateMnemonic(mnemonic) {
		return "", customerrors.ErrInvalidMnemonic
	}
	seed := utils.MnemonicToSeed(mnemonic)
	user, err := s.repo.GetBySeed(ctx, seed)
	if err != nil {
		return "", err
	}
	return utils.EncodeJWT(user.ID)
}

func (s *service) GetByID(ctx context.Context, id primitive.ObjectID) (*models.User, error) {
	return s.repo.GetByID(ctx, id)
}
