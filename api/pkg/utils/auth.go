package utils

import (
	"backend/api/pkg/config"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JWTClaims struct {
	UserId string `json:"userId"`
	jwt.StandardClaims
}

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

func EncodeJWT(userId primitive.ObjectID) (string, error) {
	claims := JWTClaims{
		UserId:         userId.Hex(),
		StandardClaims: jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.S.JwtSecret)
}

func GetUserIdFromToken(tokenString string) (primitive.ObjectID, error) {
	claims, err := DecodeJWT(tokenString)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return primitive.ObjectIDFromHex(claims.UserId)
}
