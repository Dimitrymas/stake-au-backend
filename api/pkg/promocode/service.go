package promocode

import (
	activationPkg "backend/api/pkg/activation"
	"backend/api/pkg/customerrors"
	"backend/api/pkg/models"
	"backend/api/pkg/utils"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
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
		log.Printf("promocode %s already exists", name)
		return promoObj.ID, customerrors.ErrPromoCodeExists
	} else if !errors.Is(err, customerrors.ErrPromoCodeNotFound) {
		log.Printf("create promocode: failed to get by name: %v", err)
		return primitive.NilObjectID, err
	}
	id, err := s.repository.Create(ctx, name, value, description)
	if err != nil {
		log.Printf("create promocode: repository error: %v", err)
		return primitive.NilObjectID, err
	}
	log.Printf("promocode created: %s", id.Hex())
	return id, nil
}

func (s *service) GetAll(ctx context.Context) ([]*models.PromoCode, error) {
	activations, err := s.activationService.GetAll(ctx)
	if err != nil {
		log.Printf("get promocodes: failed to load activations: %v", err)
		return nil, err
	}
	var ids []primitive.ObjectID
	for _, activation := range activations {
		ids = append(ids, activation.PromoCodeID)
	}
	ids = utils.RemoveDuplicates(ids)

	promoCodes, err := s.repository.GetByIDs(ctx, ids)
	if err != nil {
		log.Printf("get promocodes: repository error: %v", err)
		return nil, err
	}
	log.Printf("retrieved %d promocodes", len(promoCodes))
	return promoCodes, nil
}

func (s *service) GetByID(ctx context.Context, id primitive.ObjectID) (*models.PromoCode, error) {
	promo, err := s.repository.GetByID(ctx, id)
	if err != nil {
		log.Printf("get promocode by id %s: %v", id.Hex(), err)
		return nil, err
	}
	log.Printf("retrieved promocode %s", id.Hex())
	return promo, nil
}
