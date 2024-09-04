package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Promocode struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"` // Уникальный идентификатор промокода в базе
	Name        string             `bson:"name"`          // Название промокода
	Value       float64            `bson:"value"`         // Значение промокода
	Description string             `bson:"description"`   // Описание промокода
	CreatedAt   primitive.DateTime `bson:"created_at"`    // Время создания промокода
}
