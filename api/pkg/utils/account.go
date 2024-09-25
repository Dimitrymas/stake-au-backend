package utils

import (
	"backend/api/pkg/models"
	"fmt"
)

func BuildAccountProxyString(account *models.Account) string {
	if account.ProxyIP == "" || account.ProxyPort == "" {
		return ""
	}

	if account.ProxyLogin == "" || account.ProxyPass == "" {
		return fmt.Sprintf("%s:%s", account.ProxyIP, account.ProxyPort)
	}

	return fmt.Sprintf("%s:%s@%s:%s", account.ProxyLogin, account.ProxyPass, account.ProxyIP, account.ProxyPort)
}
