package dtos

import "backend/api/pkg/models"

type Account struct {
	models.Account
	Proxy          string      `json:"proxy"`
	LastActivation *Activation `json:"last_activation"`
}
