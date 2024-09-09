package routes

import (
	"backend/api/http/handlers/activation"
	middleware "backend/api/http/middlewares"
	"github.com/gofiber/fiber/v2"
)

func ActivationRouter(router fiber.Router, handler activation.CommonHandler) fiber.Router {
	api := router.Group("/activations", middleware.AuthHandler)

	api.Post("/", handler.Create)
	api.Get("/", handler.GetAll)

	return router
}
