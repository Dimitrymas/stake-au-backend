package activation

import (
	"backend/api/http/requests/activationrequests"
	"backend/api/pkg/models"
	"backend/api/pkg/utils"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Repository interface {
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

type repository struct {
	collection *mongo.Collection
}

func NewRepository(collection *mongo.Collection) Repository {
	return &repository{
		collection: collection,
	}
}

func (r *repository) Create(
	ctx context.Context,
	accountID primitive.ObjectID,
	promoCodeID primitive.ObjectID,
	succeeded bool,
	duration time.Duration,
	error string,
) error {
	newActivation := models.Activation{
		AccountID:   accountID,
		PromoCodeID: promoCodeID,
		Succeeded:   succeeded,
		Duration:    duration,
		Error:       error,
		CreatedAt:   utils.GetDateTime(),
	}

	_, err := r.collection.InsertOne(ctx, newActivation)
	return err
}

func (r *repository) CreateMany(
	ctx context.Context,
	activations []*activationrequests.Create,
) error {
	var docs []interface{}
	for _, activation := range activations {
		newActivation := models.Activation{
			AccountID:   activation.AccountID,
			PromoCodeID: activation.PromoCodeID,
			Succeeded:   activation.Succeeded,
			Duration:    activation.Duration,
			Error:       activation.Error,
			CreatedAt:   utils.GetDateTime(),
		}
		docs = append(docs, newActivation)
	}

	_, err := r.collection.InsertMany(ctx, docs)
	return err
}

func (r *repository) GetAll(ctx context.Context) ([]*models.Activation, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer utils.CloseCursor(cursor, ctx)

	var activations []*models.Activation
	if err = cursor.All(ctx, &activations); err != nil {
		return nil, err
	}

	return activations, nil
}
