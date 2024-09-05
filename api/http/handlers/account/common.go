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
}

type commonHandler struct {
	accountService accountPkg.Service
	userService    userPkg.Service
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
		data.Token,
		data.ProxyType,
		data.ProxyLogin,
		data.ProxyPass,
		data.ProxyIP,
		data.ProxyPort,
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
