package routes

import (
	"backend/api/http/handlers/user"
	middleware "backend/api/http/middlewares"
	"github.com/gofiber/fiber/v2"
)

func UserRouter(router fiber.Router, handler user.CommonHandler) fiber.Router {
	api := router.Group("/user")

	api.Post("/login", handler.Login)
	api.Post("/register", handler.Register)
	api.Get("/me", middleware.AuthHandler, handler.Me)
	api.Get("/mnemonic", handler.GenerateMnemonic)

	return router
}
