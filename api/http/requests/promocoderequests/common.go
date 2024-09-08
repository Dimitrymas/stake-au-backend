package promocoderequests

import "go.mongodb.org/mongo-driver/bson/primitive"

type Create struct {
	Name        string  `json:"name" validate:"required,min=2,max=20"`
	Value       float64 `json:"value" validate:"omitempty,min=0,max=1000"`
	Description string  `json:"description" validate:"omitempty,min=2,max=200"`
}

type GetByIDs struct {
	IDs []primitive.ObjectID `json:"ids" validate:"required"`
}
