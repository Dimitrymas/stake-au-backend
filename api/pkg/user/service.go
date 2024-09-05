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
	Register(ctx context.Context, login string, password string) (string, error)
	Login(ctx context.Context, login string, password string) (string, error)
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

func (s *service) Register(ctx context.Context, login string, password string) (string, error) {
	_, err := s.repo.GetByLogin(ctx, login)
	if err == nil {
		return "", customerrors.ErrUserAlreadyExists
	} else if !errors.Is(err, customerrors.ErrUserNotFound) {
		return "", err
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return "", err
	}

	userID, err := s.repo.Register(ctx, login, hashedPassword)
	if err != nil {
		return "", err
	}

	return utils.EncodeJWT(userID)
}

func (s *service) Login(ctx context.Context, login string, password string) (string, error) {
	user, err := s.repo.GetByLogin(ctx, login)
	if err != nil {
		return "", err
	}

	valid := utils.CheckPasswordHash(password, user.HashedPassword)
	if !valid {
		return "", customerrors.ErrPasswordIncorrect
	}

	return utils.EncodeJWT(user.ID)
}

func (s *service) GetByID(ctx context.Context, id primitive.ObjectID) (*models.User, error) {
	return s.repo.GetByID(ctx, id)
}
