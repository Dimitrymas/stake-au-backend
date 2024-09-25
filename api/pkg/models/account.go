package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Account struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"` // Уникальный идентификатор пользователя в базе
	UserID     primitive.ObjectID `bson:"user_id"`       // Уникальный идентификатор пользователя в базе
	Name       string             `bson:"name"`
	Token      string             `bson:"token"`       // Токен пользователя
	ProxyType  string             `bson:"proxy_type"`  // Тип прокси
	ProxyLogin string             `bson:"proxy_login"` // Логин прокси
	ProxyPass  string             `bson:"proxy_pass"`  // Пароль прокси
	ProxyIP    string             `bson:"proxy_ip"`    // IP прокси
	ProxyPort  string             `bson:"proxy_port"`  // Порт прокси
	CreatedAt  primitive.DateTime `bson:"created_at"`  // Время создания пользователя
}
