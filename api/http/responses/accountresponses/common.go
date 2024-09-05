package accountresponses

import "github.com/gofiber/fiber/v2"

func Created() *fiber.Map {
	return &fiber.Map{
		"message": "Account created",
	}
}

func AccountsLimit() *fiber.Map {
	return &fiber.Map{
		"message": "Accounts limit reached",
	}
}

func SubNotActive() *fiber.Map {
	return &fiber.Map{
		"message": "Subscription is not active",
	}
}
