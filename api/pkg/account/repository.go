package account

import (
	"backend/api/http/requests/accountrequests"
	"backend/api/pkg/customerrors"
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
		account *accountrequests.Create,
	) error
	GetByUserID(ctx context.Context, userID primitive.ObjectID) ([]*models.Account, error)
	CountByUserID(ctx context.Context, userID primitive.ObjectID) (int, error)
	CreateMany(
		ctx context.Context,
		userID primitive.ObjectID,
		accounts []*accountrequests.Create,
	) error
	Edit(ctx context.Context, accountID primitive.ObjectID, account *accountrequests.Edit) error
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
	requestData *accountrequests.Create,
) error {
	account := models.Account{
		UserID:     userID,
		Name:       requestData.Name,
		Token:      requestData.Token,
		ProxyType:  requestData.ProxyType,
		ProxyLogin: requestData.ProxyLogin,
		ProxyPass:  requestData.ProxyPass,
		ProxyIP:    requestData.ProxyIP,
		ProxyPort:  requestData.ProxyPort,
		CreatedAt:  utils.GetDateTime(),
	}
	// Выполняем вставку нового пользователя
	_, err := r.collection.InsertOne(ctx, account)

	return err
}

func (r *repository) GetByUserID(ctx context.Context, userID primitive.ObjectID) ([]*models.Account, error) {
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

func (r *repository) CreateMany(
	ctx context.Context,
	userID primitive.ObjectID,
	accounts []*accountrequests.Create,
) error {
	var accountsData []interface{}
	for _, account := range accounts {
		accountsData = append(accountsData, models.Account{
			UserID:     userID,
			Name:       account.Name,
			Token:      account.Token,
			ProxyType:  account.ProxyType,
			ProxyLogin: account.ProxyLogin,
			ProxyPass:  account.ProxyPass,
			ProxyIP:    account.ProxyIP,
			ProxyPort:  account.ProxyPort,
			CreatedAt:  utils.GetDateTime(),
		})
	}
	// Выполняем вставку новых аккаунтов
	_, err := r.collection.InsertMany(ctx, accountsData)

	return err
}

func (r *repository) Edit(ctx context.Context, userID primitive.ObjectID, account *accountrequests.Edit) error {
	filter := bson.M{"_id": account.ID, "user_id": userID}
	update := bson.M{
		"$set": bson.M{
			"token":       account.Token,
			"proxy_type":  account.ProxyType,
			"proxy_login": account.ProxyLogin,
			"proxy_pass":  account.ProxyPass,
			"proxy_ip":    account.ProxyIP,
			"proxy_port":  account.ProxyPort,
		},
	}
	res, err := r.collection.UpdateOne(ctx, filter, update)

	if err != nil {
		return err
	}

	if res.MatchedCount == 0 {
		return customerrors.ErrAccountNotFound
	}

	return nil
}
