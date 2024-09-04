package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`   // Уникальный идентификатор пользователя в базе
	Login          string             `bson:"login"`           // Логин пользователя
	HashedPassword string             `bson:"hashed_password"` // Хэш пароля пользователя
	SubStart       primitive.DateTime `bson:"sub_start"`       // Начало подписки
	SubEnd         primitive.DateTime `bson:"sub_end"`         // Конец подписки
	MaxAccounts    int                `bson:"max_accounts"`    // Максимальное количество аккаунтов
	CreatedAt      primitive.DateTime `bson:"created_at"`      // Время создания пользователя
}
