package routes

import (
	"backend/api/http/handlers/activation"
	"github.com/gofiber/fiber/v2"
)

func ActivationRouter(router fiber.Router, handler activation.CommonHandler) fiber.Router {
	api := router.Group("/activations")

	api.Post("/", handler.Create)
	api.Get("/", handler.GetAll)

	return router
}
