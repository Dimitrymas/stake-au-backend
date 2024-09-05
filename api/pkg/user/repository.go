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
)

type Repository interface {
	Register(ctx context.Context, login string, hashedPassword string) (primitive.ObjectID, error)
	GetByLogin(ctx context.Context, login string) (*models.User, error)
	GetByID(ctx context.Context, id primitive.ObjectID) (*models.User, error)
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
	login string,
	hashedPassword string,
) (primitive.ObjectID, error) {
	user := models.User{
		Login:          login,
		HashedPassword: hashedPassword,
		SubStart:       0,
		SubEnd:         0,
		MaxAccounts:    0,
		CreatedAt:      utils.GetDateTime(),
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

func (r *repository) GetByLogin(ctx context.Context, login string) (*models.User, error) {
	var user models.User
	filter := bson.M{"login": login}
	err := r.collection.FindOne(ctx, filter).Decode(&user)

	// Обработка ошибки, если документ не найден
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, customerrors.ErrUserNotFound
	}

	// Возвращаем ошибку, если возникла другая ошибка
	if err != nil {
		return nil, err
	}

	// Возвращаем пользователя и nil в случае успеха
	return &user, nil
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
