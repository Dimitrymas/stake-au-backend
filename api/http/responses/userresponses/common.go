package userresponses

import (
	"backend/api/pkg/models"
	"backend/api/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func Auth(token string) fiber.Map {
	return fiber.Map{
		"token": token,
	}
}

func InvalidCredentials() fiber.Map {
	return fiber.Map{
		"error": "Invalid credentials",
	}
}

func LoginAlreadyExists() fiber.Map {
	return fiber.Map{
		"error": "User with this login already exists",
	}
}

func Me(userID *models.User) fiber.Map {
	data := fiber.Map{
		"id":          userID.ID.Hex(),
		"login":       userID.Login,
		"subStart":    userID.SubStart,
		"subEnd":      userID.SubEnd,
		"maxAccounts": userID.MaxAccounts,
	}
	sign, errResponse := utils.SignData(data)
	if errResponse != nil {
		return errResponse
	}
	data["sign"] = sign
	return data
}
