package user

import (
	"backend/api/pkg/customerrors"
	"backend/api/pkg/utils"
	"context"
)

type Service interface {
	Register(ctx context.Context, login string, password string) (string, error)
	Login(ctx context.Context, login string, password string) (string, error)
}

type service struct {
	repository Repository
}

func NewService(
	repository Repository,
) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) Register(ctx context.Context, login string, password string) (string, error) {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return "", err
	}

	userID, err := s.repository.Register(ctx, login, hashedPassword)
	if err != nil {
		return "", err
	}

	return utils.EncodeJWT(userID)
}

func (s *service) Login(ctx context.Context, login string, password string) (string, error) {
	user, err := s.repository.GetByLogin(login)
	if err != nil {
		return "", err
	}

	valid := utils.CheckPasswordHash(ctx, password, user.HashedPassword)
	if !valid {
		return "", customerrors.ErrPasswordIncorrect
	}

	return utils.EncodeJWT(user.ID)
}
