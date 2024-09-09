package utils

import (
	"backend/api/pkg/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/tyler-smith/go-bip39"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type JWTClaims struct {
	UserId string `json:"userId"`
	jwt.StandardClaims
}

// DecodeJWT Функция для декодирования JWT
func DecodeJWT(tokenString string) (*JWTClaims, error) {
	var claims JWTClaims
	_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return config.S.JwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	return &claims, nil
}

// EncodeJWT Функция для кодирования JWT
func EncodeJWT(userId primitive.ObjectID) (string, error) {
	claims := JWTClaims{
		UserId:         userId.Hex(),
		StandardClaims: jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.S.JwtSecret)
}

// GetUserIdFromToken Функция для получения ID пользователя из JWT
func GetUserIdFromToken(tokenString string) (primitive.ObjectID, error) {
	claims, err := DecodeJWT(tokenString)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return primitive.ObjectIDFromHex(claims.UserId)
}

// HashPassword Функция для хэширования пароля
func HashPassword(password string) (string, error) {
	// Хэшируем пароль с использованием bcrypt и параметра cost
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// CheckPasswordHash Функция для проверки пароля
func CheckPasswordHash(password, hash string) bool {
	// Сравниваем пароль с хэшем
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GenerateMnemonic Функция для генерации мнемонической фразы
func GenerateMnemonic() ([]string, error) {
	// Генерируем мнемоническую фразу
	entropy, err := bip39.NewEntropy(128)
	if err != nil {
		return nil, err
	}
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return nil, err
	}
	return strings.Split(mnemonic, " "), nil
}

// MnemonicToSeed Функция для преобразования мнемонической фразы в seed
func MnemonicToSeed(mnemonic []string) string {
	return string(bip39.NewSeed(strings.Join(mnemonic, " "), ""))
}

func ValidateMnemonic(mnemonic []string) bool {
	return bip39.IsMnemonicValid(strings.Join(mnemonic, " "))
}
