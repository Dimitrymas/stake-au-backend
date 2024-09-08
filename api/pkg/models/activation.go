package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Activation struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"` // Уникальный идентификатор пользователя в базе
	AccountID   primitive.ObjectID `bson:"account_id"`    // ID аккаунта
	PromoCodeID primitive.ObjectID `bson:"promocode_id"`  // ID промокода
	Succeeded   bool               `bson:"succeeded"`     // Промокод успешно применен
	Duration    time.Duration      `bson:"duration"`      // Время активации промокода
	Error       string             `bson:"error"`         // Ошибка при активации промокода
	CreatedAt   primitive.DateTime `bson:"created_at"`    // Время создания промокода
}
