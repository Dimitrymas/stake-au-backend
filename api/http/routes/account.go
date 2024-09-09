package routes

import (
	"backend/api/http/handlers/account"
	middleware "backend/api/http/middlewares"
	"github.com/gofiber/fiber/v2"
)

func AccountRouter(router fiber.Router, handler account.CommonHandler) fiber.Router {
	api := router.Group("/accounts", middleware.AuthHandler) // Использовать множественное число для ресурсов

	api.Get("/", handler.GetAll)
	api.Post("/", handler.Create)         // Создание ресурса на тот же путь
	api.Post("/bulk", handler.CreateMany) // Явное указание для массового создания
	api.Patch("/", handler.Edit)          // Привязать изменение к конкретному ID

	return router
}
