package user

import (
	"backend/api/pkg/customerrors"
	"backend/api/pkg/models"
	"backend/api/pkg/utils"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Repository interface {
	Register(ctx context.Context, seed string, privateKey string, publicKey string) (primitive.ObjectID, error)
	GetByID(ctx context.Context, id primitive.ObjectID) (*models.User, error)
	GetBySeed(ctx context.Context, seed string) (*models.User, error)
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
	ctx context.Context,
	seed string,
	privateKey string,
	publicKey string,
) (primitive.ObjectID, error) {
	user := models.User{
		Seed:        seed,
		SubStart:    0,
		SubEnd:      primitive.NewDateTimeFromTime(time.Now().AddDate(0, 1, 0)),
		MaxAccounts: 3,
		PrivateKey:  privateKey,
		PublicKey:   publicKey,
		CreatedAt:   utils.GetDateTime(),
	}
	// Выполняем вставку нового пользователя
	result, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return primitive.NilObjectID, err
	}

	objectID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return primitive.NilObjectID, err
	}

	return objectID, nil
}

func (r *repository) GetByID(ctx context.Context, id primitive.ObjectID) (*models.User, error) {
	var user models.User
	filter := bson.M{"_id": id}
	err := r.collection.FindOne(ctx, filter).Decode(&user)

	// Обработка ошибки, если документ не найден
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, customerrors.ErrUserNotFound
	}
	return &user, err
}

func (r *repository) GetBySeed(ctx context.Context, seed string) (*models.User, error) {
	var user models.User
	filter := bson.M{"seed": seed}
	err := r.collection.FindOne(ctx, filter).Decode(&user)

	// Обработка ошибки, если документ не найден
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, customerrors.ErrUserNotFound
	}
	return &user, err
}
