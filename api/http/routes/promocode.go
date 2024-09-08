package routes

import (
	"backend/api/http/handlers/promocode"
	"github.com/gofiber/fiber/v2"
)

func PromoCodeRouter(router fiber.Router, handler promocode.CommonHandler) fiber.Router {
	api := router.Group("/promocodes")

	api.Get("/", handler.GetAll)
	api.Post("/", handler.Create)

	return router
}
