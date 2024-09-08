package activation

import (
	"backend/api/http/requests/activationrequests"
	"backend/api/http/responses/activationresponses"
	validation "backend/api/http/validator"
	activationPkg "backend/api/pkg/activation"
	"github.com/gofiber/fiber/v2"
)

type CommonHandler interface {
	Create(ctx *fiber.Ctx) error
	GetAll(ctx *fiber.Ctx) error
}

type commonHandler struct {
	service activationPkg.Service
}

func NewCommonHandler(service activationPkg.Service) CommonHandler {
	return &commonHandler{
		service: service,
	}
}

func (h *commonHandler) Create(ctx *fiber.Ctx) error {
	data, validationErr := validation.ParseAndValidate[activationrequests.CreateMany](ctx)
	if validationErr {
		return nil
	}
	err := h.service.CreateMany(ctx.Context(), data.Activations)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusNoContent).Send(nil)
}

func (h *commonHandler) GetAll(ctx *fiber.Ctx) error {
	activations, err := h.service.GetAll(ctx.Context())
	if err != nil {
		return err
	}
	return ctx.JSON(activationresponses.GetAll(activations))
}
