package promocoderesponses

import (
	"backend/api/pkg/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Created(promoID primitive.ObjectID) fiber.Map {
	return fiber.Map{
		"promoID": promoID.Hex(),
	}
}

func Get(promoCode *models.PromoCode) fiber.Map {
	return fiber.Map{
		"id":          promoCode.ID.Hex(),
		"name":        promoCode.Name,
		"value":       promoCode.Value,
		"description": promoCode.Description,
		"createdAt":   promoCode.CreatedAt,
	}
}

func GetAll(promoCodes []*models.PromoCode) fiber.Map {
	var response []fiber.Map
	for _, promoCode := range promoCodes {
		response = append(response, Get(promoCode))
	}
	return fiber.Map{
		"promoCodes": response,
	}
}
