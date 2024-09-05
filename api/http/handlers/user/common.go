package user

import (
	"backend/api/http/requests/userrequests"
	"backend/api/http/responses/userresponses"
	validation "backend/api/http/validator"
	"backend/api/pkg/customerrors"
	userPkg "backend/api/pkg/user"
	"errors"
	"github.com/gofiber/fiber/v2"
)

type CommonHandler interface {
}

type commonHandler struct {
	service userPkg.Service
}

func (h *commonHandler) Register(ctx *fiber.Ctx) error {
	data, validationErr := validation.ParseAndValidate[userrequests.Register](ctx)
	if validationErr {
		return nil
	}

	token, err := h.service.Register(ctx.Context(), data.Login, data.Password)
	switch {
	case errors.Is(err, nil):
		return ctx.JSON(userresponses.AuthData(token))
	case errors.Is(err, customerrors.ErrUserAlreadyExists):
		return ctx.Status(fiber.StatusConflict).JSON(userresponses.LoginAlreadyExists())
	default:
		return err
	}

}

func (h *commonHandler) Login(ctx *fiber.Ctx) error {
	data, validationErr := validation.ParseAndValidate[userrequests.Login](ctx)
	if validationErr {
		return nil
	}

	token, err := h.service.Login(ctx.Context(), data.Login, data.Password)
	switch {
	case errors.Is(err, nil):
		return ctx.JSON(userresponses.AuthData(token))
	case errors.Is(err, customerrors.ErrUserNotFound):
		return ctx.Status(fiber.StatusUnauthorized).JSON(userresponses.InvalidCredentials())
	case errors.Is(err, customerrors.ErrPasswordIncorrect):
		return ctx.Status(fiber.StatusUnauthorized).JSON(userresponses.InvalidCredentials())
	default:
		return err
	}
}
