package middleware

import (
	"backend/api/http/responses"
	"github.com/gofiber/fiber/v2"
)

// ErrorHandler is a middleware that handles errors and returns a 500 Internal Server Error status
func ErrorHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// Call the next handler in the chain
		err := ctx.Next()
		if err != nil {
			// Return a 500 Internal Server Error response
			return ctx.Status(fiber.StatusInternalServerError).JSON(responses.InternalError())
		}
		return nil
	}
}
