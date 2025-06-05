package middleware

import (
	"backend/api/http/responses"
	"backend/api/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"log"
)

func AuthHandler(ctx *fiber.Ctx) error {
	// Get the value of the Authorization header
	token := ctx.Get("Authorization")
	// If the Authorization header is missing
	if len(token) < 7 || token[:7] != "Bearer " {
		log.Println("authorization header missing or invalid")
		return ctx.Status(fiber.StatusUnauthorized).JSON(responses.Unauthorized())
	}

	token = token[7:]

	// Verify the token
	userID, err := utils.GetUserIdFromToken(token)

	if err != nil {
		log.Printf("invalid token: %v", err)
		return ctx.Status(fiber.StatusUnauthorized).JSON(responses.Unauthorized())
	}

	// Store the user ID in the locals
	ctx.Locals("userID", userID)
	return ctx.Next()
}
