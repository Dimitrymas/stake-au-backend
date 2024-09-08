package activationresponses

import (
	"backend/api/pkg/models"
	"github.com/gofiber/fiber/v2"
)

func Created() fiber.Map {
	return fiber.Map{
		"message": "activation created",
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
