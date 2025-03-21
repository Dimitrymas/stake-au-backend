package accountresponses

import (
	"backend/api/http/responses/activationresponses"
	"backend/api/pkg/dtos"
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
		"error": "Accounts limit reached",
	}
}

func SubNotActive() fiber.Map {
	return fiber.Map{
		"error": "Subscription is not active",
	}
}

func GetAccount(account *dtos.Account) fiber.Map {
	return fiber.Map{
		"id":             account.ID.Hex(),
		"token":          account.Token,
		"proxyType":      account.ProxyType,
		"proxyLogin":     account.ProxyLogin,
		"proxyPass":      account.ProxyPass,
		"proxyIP":        account.ProxyIP,
		"proxy_port":     account.ProxyPort,
		"proxy":          account.Proxy,
		"lastActivation": activationresponses.GetWithPromoCode(account.LastActivation),
		"createdAt":      account.CreatedAt,
	}
}

func Get(accounts []*dtos.Account, privateKeyEnc string) fiber.Map {
	result := make([]fiber.Map, 0, len(accounts))
	for _, account := range accounts {
		result = append(result, GetAccount(account))
	}
	data := fiber.Map{
		"accounts": result,
	}
	return utils.SignData(data, privateKeyEnc)
}

func NotFound() fiber.Map {
	return fiber.Map{
		"error": "Account not found",
	}
}

func Edited() fiber.Map {
	return fiber.Map{
		"message": "Account edited",
	}
}
