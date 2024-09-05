package validation

import (
	"backend/api/http/responses"
	"github.com/gofiber/fiber/v2"
)

// ParseAndValidate parses the request body and validates it, returning the parsed data or an error
func ParseAndValidate[T any](ctx *fiber.Ctx) (*T, bool) {
	var data T
	if err := ctx.BodyParser(&data); err != nil {
		//return data, ctx.Status(fiber.StatusBadRequest).JSON(responses.BadRequest())
		err = ctx.Status(fiber.StatusBadRequest).JSON(responses.BadRequest())
		return nil, true
	}
	if err := ValidateStruct(&data); err != nil {
		err = ctx.Status(fiber.StatusBadRequest).JSON(
			responses.BadRequestWithErrors(
				HandleValidationError(err),
			),
		)
		return nil, true
	}
	return &data, false
}
