package promocode

import (
	"backend/api/http/requests/accountrequests"
	validation "backend/api/http/validator"
	userPkg "backend/api/pkg/user"
	"github.com/gofiber/fiber/v2"
)

type CommonHandler interface {
}

type commonHandler struct {
	service userPkg.Service
}

func NewCommonHandler() CommonHandler {
	return &commonHandler{}
}

func (h *commonHandler) Create(ctx *fiber.Ctx) error {
	data, validationErr := validation.ParseAndValidate[accountrequests.Create](ctx)
	if validationErr {
		return nil
	}

}
