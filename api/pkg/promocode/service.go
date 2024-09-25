package promocode

import (
	activationPkg "backend/api/pkg/activation"
	"backend/api/pkg/customerrors"
	"backend/api/pkg/models"
	"backend/api/pkg/utils"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
)

type Service interface {
	Create(ctx context.Context, name string, value float64, description string) (primitive.ObjectID, error)
	GetAll(ctx context.Context) ([]*models.PromoCode, error)
	GetByID(ctx context.Context, id primitive.ObjectID) (*models.PromoCode, error)
}

type service struct {
	repository        Repository
	activationService activationPkg.Service
}

func NewService(repository Repository, activationService activationPkg.Service) Service {
	return &service{
		repository:        repository,
		activationService: activationService,
	}
}

func (s *service) Create(ctx context.Context, name string, value float64, description string) (primitive.ObjectID, error) {
	name = strings.ToUpper(name)
	promoObj, err := s.repository.GetByName(ctx, name)
	if err == nil {
		return promoObj.ID, customerrors.ErrPromoCodeExists
	} else if !errors.Is(err, customerrors.ErrPromoCodeNotFound) {
		return primitive.NilObjectID, err
	}
	return s.repository.Create(ctx, name, value, description)
}

func (s *service) GetAll(ctx context.Context) ([]*models.PromoCode, error) {
	activations, err := s.activationService.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	var ids []primitive.ObjectID
	for _, activation := range activations {
		ids = append(ids, activation.PromoCodeID)
	}
	ids = utils.RemoveDuplicates(ids)

	return s.repository.GetByIDs(ctx, ids)
}

func (s *service) GetByID(ctx context.Context, id primitive.ObjectID) (*models.PromoCode, error) {
	return s.repository.GetByID(ctx, id)
}
