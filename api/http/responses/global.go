package responses

import "github.com/gofiber/fiber/v2"

func BadRequest() *fiber.Map {
	return &fiber.Map{
		"error": "Bad Request",
	}
}

func BadRequestWithErrors(errors map[string]string) *fiber.Map {
	return &fiber.Map{
		"error":  "Bad Request",
		"errors": errors,
	}
}

func InternalError() *fiber.Map {
	return &fiber.Map{
		"error": "Internal Server Error",
	}
}

func Unauthorized() *fiber.Map {
	return &fiber.Map{
		"error": "Unauthorized",
	}
}
