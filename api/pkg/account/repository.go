package account

import (
	"backend/api/pkg/models"
	"backend/api/pkg/utils"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	Create(
		ctx context.Context,
		userID primitive.ObjectID,
		token string,
		proxyType string,
		proxyLogin string,
		proxyPass string,
		proxyIP string,
		proxyPort string,
	) error
	GetAllByUserID(ctx context.Context, userID primitive.ObjectID) ([]*models.Account, error)
	CountByUserID(ctx context.Context, userID primitive.ObjectID) (int, error)
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
	userID primitive.ObjectID,
	token string,
	proxyType string,
	proxyLogin string,
	proxyPass string,
	proxyIP string,
	proxyPort string,
) error {
	account := models.Account{
		UserID:     userID,
		Token:      token,
		ProxyType:  proxyType,
		ProxyLogin: proxyLogin,
		ProxyPass:  proxyPass,
		ProxyIP:    proxyIP,
		ProxyPort:  proxyPort,
		CreatedAt:  utils.GetDateTime(),
	}
	// Выполняем вставку нового пользователя
	_, err := r.collection.InsertOne(ctx, account)

	return err
}

func (r *repository) GetAllByUserID(ctx context.Context, userID primitive.ObjectID) ([]*models.Account, error) {
	filter := bson.M{"user_id": userID}
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var accounts []*models.Account
	defer utils.CloseCursor(cursor, ctx)
	err = cursor.All(ctx, &accounts)
	return accounts, err
}

func (r *repository) CountByUserID(ctx context.Context, userID primitive.ObjectID) (int, error) {
	filter := bson.M{"user_id": userID}
	count, err := r.collection.CountDocuments(ctx, filter)
	return int(count), err
}
