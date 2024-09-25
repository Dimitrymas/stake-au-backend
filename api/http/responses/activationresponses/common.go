package activationresponses

import (
	"backend/api/http/responses/promocoderesponses"
	"backend/api/pkg/dtos"
	"backend/api/pkg/models"
	"github.com/gofiber/fiber/v2"
)

func GetWithPromoCode(activation *dtos.Activation) fiber.Map {
	if activation == nil {
		return nil
	}
	return fiber.Map{
		"id":        activation.ID.Hex(),
		"promoCode": promocoderesponses.Get(activation.PromoCode),
		"succeeded": activation.Succeeded,
		"duration":  activation.Duration,
		"error":     activation.Error,
		"createdAt": activation.CreatedAt,
	}
}

func Get(activation *models.Activation) fiber.Map {
	return fiber.Map{
		"id":          activation.ID.Hex(),
		"promoCodeID": activation.PromoCodeID.Hex(),
		"accountID":   activation.AccountID.Hex(),
		"succeeded":   activation.Succeeded,
		"duration":    activation.Duration,
		"error":       activation.Error,
		"createdAt":   activation.CreatedAt,
	}
}

func GetAll(activations []*models.Activation) fiber.Map {
	var response []fiber.Map
	for _, activation := range activations {
		response = append(response, Get(activation))
	}
	return fiber.Map{
		"activations": response,
	}
}
