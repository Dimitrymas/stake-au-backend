package user

import (
	"backend/api/http/requests/userrequests"
	"backend/api/http/responses/userresponses"
	validation "backend/api/http/validator"
	"backend/api/pkg/customerrors"
	userPkg "backend/api/pkg/user"
	"backend/api/pkg/utils"
	"errors"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommonHandler interface {
	GenerateMnemonic(ctx *fiber.Ctx) error
	Register(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
	Me(ctx *fiber.Ctx) error
}

type commonHandler struct {
	service userPkg.Service
}

func NewCommonHandler(
	service userPkg.Service,
) CommonHandler {
	return &commonHandler{
		service: service,
	}
}

func (h *commonHandler) GenerateMnemonic(ctx *fiber.Ctx) error {
	if ctx.Locals("userID") != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(userresponses.AlreadyAuthenticated())
	}
	mnemonic, err := utils.GenerateMnemonic()
	if err != nil {
		return err
	}
	return ctx.JSON(userresponses.Mnemonic(mnemonic))
}

func (h *commonHandler) Register(ctx *fiber.Ctx) error {
	data, validationErr := validation.ParseAndValidate[userrequests.Register](ctx)
	if validationErr {
		return nil
	}

	token, publicKey, privateKey, err := h.service.Register(ctx.Context(), data.Mnemonic)
	switch {
	case err == nil:
		return ctx.JSON(userresponses.Auth(token, publicKey, privateKey))
	case errors.Is(err, customerrors.ErrUserAlreadyExists):
		return ctx.Status(fiber.StatusConflict).JSON(userresponses.MnemonicAlreadyExists())
	case errors.Is(err, customerrors.ErrInvalidMnemonic):
		return ctx.Status(fiber.StatusBadRequest).JSON(userresponses.InvalidMnemonic())
	default:
		return err
	}
}

func (h *commonHandler) Login(ctx *fiber.Ctx) error {
	data, validationErr := validation.ParseAndValidate[userrequests.Login](ctx)
	if validationErr {
		return nil
	}

	token, publicKey, privateKey, err := h.service.Login(ctx.Context(), data.Mnemonic)
	switch {
	case err == nil:
		return ctx.JSON(userresponses.Auth(token, publicKey, privateKey))
	case errors.Is(err, customerrors.ErrUserNotFound):
		return ctx.Status(fiber.StatusUnauthorized).JSON(userresponses.InvalidCredentials())
	case errors.Is(err, customerrors.ErrInvalidMnemonic):
		return ctx.Status(fiber.StatusUnauthorized).JSON(userresponses.InvalidCredentials())
	case errors.Is(err, customerrors.ErrPasswordIncorrect):
		return ctx.Status(fiber.StatusUnauthorized).JSON(userresponses.InvalidCredentials())
	default:
		return err
	}
}

func (h *commonHandler) Me(ctx *fiber.Ctx) error {
	userID := ctx.Locals("userID").(primitive.ObjectID)

	user, err := h.service.GetByID(ctx.Context(), userID)
	if err != nil {
		return err
	}
	return ctx.JSON(userresponses.Me(user))
}
