package accountresponses

import (
	"backend/api/pkg/models"
	"backend/api/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func Created() fiber.Map {
	return fiber.Map{
		"message": "Account created",
	}
}

func CreatedMany(warning string) fiber.Map {
	return fiber.Map{
		"message": "Accounts created",
		"warning": warning,
	}
}

func AccountsLimit() fiber.Map {
	return fiber.Map{
		"message": "Accounts limit reached",
	}
}

func SubNotActive() fiber.Map {
	return fiber.Map{
		"message": "Subscription is not active",
	}
}

func GetAccount(account *models.Account) fiber.Map {
	return fiber.Map{
		"id":          account.ID.Hex(),
		"token":       account.Token,
		"proxy_type":  account.ProxyType,
		"proxy_login": account.ProxyLogin,
		"proxy_pass":  account.ProxyPass,
		"proxy_ip":    account.ProxyIP,
		"proxy_port":  account.ProxyPort,
		"created_at":  account.CreatedAt,
	}
}

func Get(accounts []*models.Account) fiber.Map {
	result := make([]fiber.Map, 0, len(accounts))
	for _, account := range accounts {
		result = append(result, GetAccount(account))
	}
	sign, errResponse := utils.SignData(result)
	if errResponse != nil {
		return errResponse
	}
	return fiber.Map{
		"accounts": result,
		"sign":     sign,
	}
}
