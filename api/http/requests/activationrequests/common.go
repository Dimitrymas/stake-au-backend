package activationrequests

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Create struct {
	AccountID   primitive.ObjectID `json:"account_id" validate:"required"`
	PromoCodeID primitive.ObjectID `json:"promocode_id" validate:"required"`
	Succeeded   bool               `json:"succeeded" validate:"required"`
	Duration    time.Duration      `json:"duration" validate:"required"`
	Error       string             `json:"error" validate:"required"`
}

type CreateMany struct {
	Activations []*Create `json:"activations" validate:"required,dive"`
}
