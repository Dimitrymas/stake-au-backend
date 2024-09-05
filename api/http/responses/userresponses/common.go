package userresponses

import (
	"github.com/gofiber/fiber/v2"
)

func AuthData(token string) *fiber.Map {
	return &fiber.Map{
		"token": token,
	}
}

func InvalidCredentials() *fiber.Map {
	return &fiber.Map{
		"error": "Invalid credentials",
	}
}

func LoginAlreadyExists() *fiber.Map {
	return &fiber.Map{
		"error": "User with this login already exists",
	}
}
