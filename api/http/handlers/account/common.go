package account

import (
	"backend/api/http/requests/accountrequests"
	"backend/api/http/responses/accountresponses"
	validation "backend/api/http/validator"
	accountPkg "backend/api/pkg/account"
	"backend/api/pkg/customerrors"
	userPkg "backend/api/pkg/user"
	"errors"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommonHandler interface {
	Create(ctx *fiber.Ctx) error
	GetAll(ctx *fiber.Ctx) error
	CreateMany(ctx *fiber.Ctx) error
	Edit(ctx *fiber.Ctx) error
}

type commonHandler struct {
	accountService accountPkg.Service
	userService    userPkg.Service
}

func NewCommonHandler(
	accountService accountPkg.Service,
	userService userPkg.Service,
) CommonHandler {
	return &commonHandler{
		accountService: accountService,
		userService:    userService,
	}
}

func (h *commonHandler) Create(ctx *fiber.Ctx) error {
	userID := ctx.Locals("userID").(primitive.ObjectID)

	data, validationErr := validation.ParseAndValidate[accountrequests.Create](ctx)
	if validationErr {
		return nil
	}

	err := h.accountService.Create(
		ctx.Context(),
		userID,
		data,
	)
	switch {
	case err == nil:
		return ctx.JSON(accountresponses.Created())
	case errors.Is(err, customerrors.ErrSubNotActive):
		return ctx.Status(fiber.StatusForbidden).JSON(accountresponses.SubNotActive())
	case errors.Is(err, customerrors.ErrAccountsLimit):
		return ctx.Status(fiber.StatusForbidden).JSON(accountresponses.AccountsLimit())
	default:
		return err
	}
}

func (h *commonHandler) GetAll(ctx *fiber.Ctx) error {
	userID := ctx.Locals("userID").(primitive.ObjectID)

	accounts, err := h.accountService.GetByUserID(ctx.Context(), userID)
	if err != nil {
		return err
	}
	return ctx.JSON(accountresponses.Get(accounts))
}

func (h *commonHandler) CreateMany(ctx *fiber.Ctx) error {
	userID := ctx.Locals("userID").(primitive.ObjectID)

	data, validationErr := validation.ParseAndValidate[accountrequests.CreateMany](ctx)
	if validationErr {
		return nil
	}

	err := h.accountService.CreateMany(
		ctx.Context(),
		userID,
		data.Accounts,
	)

	switch {
	case err == nil:
		return ctx.JSON(accountresponses.CreatedMany(""))
	case errors.Is(err, customerrors.ErrSubNotActive):
		return ctx.Status(fiber.StatusForbidden).JSON(accountresponses.SubNotActive())
	case errors.Is(err, customerrors.ErrAccountsLimit):
		return ctx.Status(fiber.StatusForbidden).JSON(accountresponses.AccountsLimit())
	case errors.Is(err, customerrors.ErrCreatePartialAccounts):
		return ctx.Status(fiber.StatusForbidden).JSON(accountresponses.CreatedMany(err.Error()))
	default:
		return err
	}
}

func (h *commonHandler) Edit(ctx *fiber.Ctx) error {
	userID := ctx.Locals("userID").(primitive.ObjectID)

	data, validationErr := validation.ParseAndValidate[accountrequests.Edit](ctx)
	if validationErr {
		return nil
	}

	err := h.accountService.Edit(
		ctx.Context(),
		userID,
		data,
	)
	switch {
	case err == nil:
		return ctx.JSON(accountresponses.Edited())
	case errors.Is(err, customerrors.ErrAccountNotFound):
		return ctx.Status(fiber.StatusNotFound).JSON(accountresponses.NotFound())
	default:
		return err
	}
}
