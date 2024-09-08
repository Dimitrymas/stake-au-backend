package activation

import (
	"backend/api/http/requests/activationrequests"
	"backend/api/pkg/models"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Service interface {
	Create(
		ctx context.Context,
		accountID primitive.ObjectID,
		promoCodeID primitive.ObjectID,
		succeeded bool,
		duration time.Duration,
		error string,
	) error
	CreateMany(
		ctx context.Context,
		activations []*activationrequests.Create,
	) error
	GetAll(ctx context.Context) ([]*models.Activation, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) Create(
	ctx context.Context,
	accountID primitive.ObjectID,
	promoCodeID primitive.ObjectID,
	succeeded bool,
	duration time.Duration,
	error string,
) error {
	return s.repo.Create(ctx, accountID, promoCodeID, succeeded, duration, error)
}

func (s *service) CreateMany(
	ctx context.Context,
	activations []*activationrequests.Create,
) error {
	return s.repo.CreateMany(ctx, activations)
}

func (s *service) GetAll(ctx context.Context) ([]*models.Activation, error) {
	return s.repo.GetAll(ctx)
}
