package routes

import (
	"backend/api/http/handlers/account"
	"backend/api/http/handlers/activation"
	"backend/api/http/handlers/promocode"
	"backend/api/http/handlers/user"
	"github.com/gofiber/fiber/v2"
)

func Router(
	app fiber.Router,
	userHandler user.CommonHandler,
	promoCodeHandler promocode.CommonHandler,
	activationHandler activation.CommonHandler,
	accountHandler account.CommonHandler,
) fiber.Router {
	router := app.Group("/api")

	UserRouter(router, userHandler)
	PromoCodeRouter(router, promoCodeHandler)
	ActivationRouter(router, activationHandler)
	AccountRouter(router, accountHandler)

	return app
}
