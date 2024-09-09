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
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func OxapaySign(callback []byte) string {
	hmacObj := hmac.New(sha512.New, []byte(config.S.OxapayMerchantApiKey))
	hmacObj.Write(callback)
	return hex.EncodeToString(hmacObj.Sum(nil))
}

// SignData Функция для подписания данных
func SignData(data fiber.Map, privateKeyEnc string) fiber.Map {
	privateKey, err := LoadPrivateKey(privateKeyEnc)
	if err != nil {
		return fiber.Map{
			"error": "Failed to load private key",
		}
	}
	// Преобразуем данные в байтовый массив
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return fiber.Map{
			"error": "Failed to marshal data",
		}
	}

	hashed := sha256.Sum256(dataBytes)

	// Подписываем хеш с использованием приватного ключа
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		return fiber.Map{
			"error": "Failed to sign data",
		}
	}

	signB64 := base64.StdEncoding.EncodeToString(signature)
	data["sign"] = signB64

	return data
}

// LoadPrivateKey Функция для загрузки приватного ключа
func LoadPrivateKey(privateKeyStr string) (*rsa.PrivateKey, error) {
	// Декодируем base64 строку
	privateKeyPem, err := base64.StdEncoding.DecodeString(privateKeyStr)
	if err != nil {
		return nil, fmt.Errorf("error decoding base64 private key: %v", err)
	}

	// Декодируем PEM блок
	block, _ := pem.Decode(privateKeyPem)
	if block == nil {
		return nil, errors.New("error decoding PEM block: PEM block is nil")
	}

	// Парсим приватный ключ PKCS1
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("error parsing PKCS1 private key: %v", err)
	}

	return privateKey, nil
}

// GenerateKeyPair Функция для генерации пары ключей
func GenerateKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}
	publicKey := &privateKey.PublicKey
	return privateKey, publicKey, nil
}

// PrivateKeyToString Функция для конвертации приватного ключа в строку (PEM формат)
func PrivateKeyToString(privateKey *rsa.PrivateKey) string {
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privatePem := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	})
	return base64.StdEncoding.EncodeToString(privatePem)
}

// PublicKeyToString Функция для конвертации публичного ключа в строку (PEM формат)
func PublicKeyToString(pubkey *rsa.PublicKey) (string, error) {
	pubKeyBytes, err := x509.MarshalPKIXPublicKey(pubkey)
	if err != nil {
		return "", err
	}
	publicPem := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubKeyBytes,
	})
	return base64.StdEncoding.EncodeToString(publicPem), nil
}
