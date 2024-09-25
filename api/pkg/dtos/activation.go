package dtos

import "backend/api/pkg/models"

type Activation struct {
	models.Activation
	PromoCode *models.PromoCode
}
