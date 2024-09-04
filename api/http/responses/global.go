package responses

import "github.com/gofiber/fiber/v2"

func GlobalBadRequest(errors *map[string]string) *fiber.Map {
	return &fiber.Map{
		"success": false,
		"error":   "Bad Request",
		"errors":  errors,
	}
}

func GlobalInternalError() *fiber.Map {
	return &fiber.Map{
		"success": false,
		"error":   "Internal Server Error",
	}
}
