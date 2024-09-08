package promocode

import (
	"backend/api/http/requests/promocoderequests"
	"backend/api/http/responses/promocoderesponses"
	validation "backend/api/http/validator"
	"backend/api/pkg/customerrors"
	promoCodePkg "backend/api/pkg/promocode"
	"errors"
	"github.com/gofiber/fiber/v2"
)

type CommonHandler interface {
	Create(ctx *fiber.Ctx) error
	GetAll(ctx *fiber.Ctx) error
}

type commonHandler struct {
	service promoCodePkg.Service
}

func NewCommonHandler(promoCodeService promoCodePkg.Service) CommonHandler {
	return &commonHandler{
		service: promoCodeService,
	}
}

func (h *commonHandler) Create(ctx *fiber.Ctx) error {
	data, validationErr := validation.ParseAndValidate[promocoderequests.Create](ctx)
	if validationErr {
		return nil
	}
	promoID, err := h.service.Create(ctx.Context(), data.Name, data.Value, data.Description)
	if err != nil && !errors.Is(err, customerrors.ErrPromoCodeExists) {
		return err
	}
	return ctx.JSON(promocoderesponses.Created(promoID))
}

func (h *commonHandler) GetAll(ctx *fiber.Ctx) error {
	promoCodes, err := h.service.GetAll(ctx.Context())
	if err != nil {
		return err
	}
	return ctx.JSON(promocoderesponses.GetAll(promoCodes))
}
