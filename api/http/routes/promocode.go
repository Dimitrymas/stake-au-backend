package routes

import (
	"backend/api/http/handlers/promocode"
	middleware "backend/api/http/middlewares"
	"github.com/gofiber/fiber/v2"
)

func PromoCodeRouter(router fiber.Router, handler promocode.CommonHandler) fiber.Router {
	api := router.Group("/promocodes", middleware.AuthHandler)

	api.Get("/", handler.GetAll)
	api.Post("/", handler.Create)

	return router
}
