package promocode

import (
	"backend/api/pkg/customerrors"
	"backend/api/pkg/models"
	"backend/api/pkg/utils"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	Create(ctx context.Context, name string, value float64, description string) (primitive.ObjectID, error)
	GetByName(ctx context.Context, name string) (*models.PromoCode, error)
	GetByIDs(ctx context.Context, ids []primitive.ObjectID) ([]*models.PromoCode, error)
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
	name string,
	value float64,
	description string,
) (primitive.ObjectID, error) {
	promoCode := models.PromoCode{
		Name:        name,
		Value:       value,
		Description: description,
		CreatedAt:   utils.GetDateTime(),
	}

	resp, err := r.collection.InsertOne(ctx, promoCode)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return resp.InsertedID.(primitive.ObjectID), err
}

func (r *repository) GetByName(ctx context.Context, name string) (*models.PromoCode, error) {
	promoCode := new(models.PromoCode)
	filter := bson.M{"name": name}
	err := r.collection.FindOne(ctx, filter).Decode(promoCode)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, customerrors.ErrPromoCodeNotFound
	}
	return promoCode, err
}

func (r *repository) GetByIDs(ctx context.Context, ids []primitive.ObjectID) ([]*models.PromoCode, error) {
	cursor, err := r.collection.Find(ctx, bson.M{"_id": bson.M{"$in": ids}})
	if err != nil {
		return nil, err
	}
	defer utils.CloseCursor(cursor, ctx)

	var promoCodes []*models.PromoCode
	if err = cursor.All(ctx, &promoCodes); err != nil {
		return nil, err
	}
	return promoCodes, nil
}
