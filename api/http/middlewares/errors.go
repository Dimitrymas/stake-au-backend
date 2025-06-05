package middleware

import (
	"backend/api/http/responses"
	"errors"
	"github.com/gofiber/fiber/v2"
	"log"
)

// ErrorHandler is a middleware that handles errors and returns a 500 Internal Server Error status
func ErrorHandler(ctx *fiber.Ctx) error {
	// Call the next handler in the chain
	err := ctx.Next()
	if err != nil {
		var cError *fiber.Error
		if errors.As(err, &cError) {
			return err
		}
		log.Printf("internal server error: %v", err)
		// Return a 500 Internal Server Error response
		return ctx.Status(fiber.StatusInternalServerError).JSON(responses.InternalError())
	}
	return nil
}
