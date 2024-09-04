package utils

import (
	"backend/api/pkg/config"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
)

func OxapaySign(callback []byte) string {
	hmacObj := hmac.New(sha512.New, []byte(config.S.OxapayMerchantApiKey))
	hmacObj.Write(callback)
	return hex.EncodeToString(hmacObj.Sum(nil))
}
