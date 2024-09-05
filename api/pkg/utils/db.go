package utils

import (
	"backend/api/pkg/config"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func DatabaseConnection() (*mongo.Database, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	url := config.S.GetDbUrl()
	log.Println("Connecting to database")

	opts := options.Client().ApplyURI(url).SetServerSelectionTimeout(5 * time.Second)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Println(err)
		cancel()
		return nil, nil, err
	}

	// Проверяем соединение с базой данных с помощью Ping
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Println("Failed to ping database:", err)
		cancel()
		return nil, nil, err
	}

	db := client.Database(config.S.DbName)
	return db, cancel, nil
}

func CloseCursor(cursor *mongo.Cursor, ctx context.Context) {
	if err := cursor.Close(ctx); err != nil {
		log.Println(err)
	}
}

func GetDateTime() primitive.DateTime {
	return primitive.NewDateTimeFromTime(time.Now())
}
