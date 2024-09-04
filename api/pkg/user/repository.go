package user

import (
	"backend/api/pkg/models"
	"backend/api/pkg/utils"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	Register(
		login string,
		hashedPassword string,
	) error
}

type repository struct {
	collection *mongo.Collection
}

func NewRepository(collection *mongo.Collection) Repository {
	return &repository{
		collection: collection,
	}
}

func (r *repository) Register(
	login string,
	hashedPassword string,
) error {
	user := models.User{
		Login:          login,
		HashedPassword: hashedPassword,
		SubStart:       0,
		SubEnd:         0,
		MaxAccounts:    0,
		CreatedAt:      utils.GetDateTime(),
	}
	_, err := r.collection.InsertOne(context.Background(), user)
	return err
}
