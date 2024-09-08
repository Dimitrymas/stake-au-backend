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

func MnemonicAlreadyExists() fiber.Map {
	return fiber.Map{
		"error": "User with this mnemonic already exists",
	}
}

func Me(userObj *models.User) fiber.Map {
	data := fiber.Map{
		"id":          userObj.ID.Hex(),
		"subStart":    userObj.SubStart,
		"subEnd":      userObj.SubEnd,
		"maxAccounts": userObj.MaxAccounts,
	}
	return utils.SignData(data)
}

func Mnemonic(mnemonic []string) fiber.Map {
	return fiber.Map{
		"mnemonic": mnemonic,
	}
}

func InvalidMnemonic() fiber.Map {
	return fiber.Map{
		"error": "Invalid mnemonic",
	}
}
