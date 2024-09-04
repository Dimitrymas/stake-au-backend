package user

import (
	"backend/api/http/requests/userrequests"
	validation "backend/api/http/validator"
	userPkg "backend/api/pkg/user"
	"github.com/gofiber/fiber/v2"
)

type CommonHandler interface {
}

type commonHandler struct {
	service userPkg.Service
}

func (h *commonHandler) Register(ctx *fiber.Ctx) error {
	var data userrequests.RegisterRequest
	err := ctx.BodyParser(&data)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(responses.GlobalBadRequest())
	}
	if err := validation.ValidateStruct(&data); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			responses.GlobalBadRequestWithDetails(validation.TranslateError(err)),
		)
	}
}
