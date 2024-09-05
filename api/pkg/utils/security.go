package utils

import (
	"backend/api/pkg/config"
	"crypto"
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"errors"
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
)

var (
	privateKey *rsa.PrivateKey
)

func OxapaySign(callback []byte) string {
	hmacObj := hmac.New(sha512.New, []byte(config.S.OxapayMerchantApiKey))
	hmacObj.Write(callback)
	return hex.EncodeToString(hmacObj.Sum(nil))
}

// Загрузка приватного ключа
func loadPrivateKey(filename string) error {
	keyData, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	block, _ := pem.Decode(keyData)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return errors.New("failed to decode PEM block containing private key")
	}

	privateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	return err
}

func init() {
	if err := loadPrivateKey("certs/private.pem"); err != nil {
		panic(err)
	}
}

// SignData Функция для подписания данных
func SignData(data interface{}) (string, fiber.Map) {
	// Преобразуем данные в байтовый массив
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return "", fiber.Map{
			"error": "Failed to marshal data",
		}
	}

	hashed := sha256.Sum256(dataBytes)

	// Подписываем хеш с использованием приватного ключа
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		return "", fiber.Map{
			"error": "Failed to sign data",
		}
	}

	// Возвращаем подпись в виде строки base64
	return base64.StdEncoding.EncodeToString(signature), nil
}
