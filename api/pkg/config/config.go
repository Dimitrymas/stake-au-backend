package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log"
)

type specification struct {
	DbHost string `envconfig:"DB_HOST"`
	DbPort string `envconfig:"DB_PORT"`
	DbName string `envconfig:"DB_NAME"`
	DbUser string `envconfig:"DB_USER"`
	DbPass string `envconfig:"DB_PASS"`

	OxapayMerchantApiKey string `envconfig:"OXAPAY_MERCHANT_API_KEY"`
	OxapayCallbackUrl    string `envconfig:"OXAPAY_CALLBACK_URL"`

	JwtSecret []byte `envconfig:"JWT_SECRET"`
}

var S specification

func (s *specification) GetDbUrl() string {
	return "mongodb://" + s.DbUser + ":" + s.DbPass + "@" + s.DbHost + ":" + s.DbPort + "/" + "admin"
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("No .env file found: %v", err)
	}
	err = envconfig.Process("", &S)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("Starting application with settings: %+v", S)
}
