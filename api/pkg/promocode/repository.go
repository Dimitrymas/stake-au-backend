package promocode

import (
	"backend/api/pkg/models"
	"backend/api/pkg/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	Create(name string, value float64, description string) error
}

type repository struct {
	collection *mongo.Collection
}

func NewRepository(collection *mongo.Collection) Repository {
	return &repository{
		collection: collection,
	}
}

func (r *repository) Create(name string, value float64, description string) error {
	promoCode := models.PromoCode{
		Name:        name,
		Value:       value,
		Description: description,
		CreatedAt:   utils.GetDateTime(),
	}

	_, err := r.collection.InsertOne(nil, promoCode)
	return err
}
