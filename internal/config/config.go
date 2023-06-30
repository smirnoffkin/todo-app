package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort             string
	PostgresURL         string
	SaltForHashPassword string
	SecretKey           string
}

func NewConfig() *Config {
	return &Config{
		AppPort:             getEnv("APP_PORT", "8000"),
		PostgresURL:         getEnv("POSTGRES_URL", ""),
		SaltForHashPassword: getEnv("SALT_FOR_HASHING_PASSWORD", ""),
		SecretKey:           getEnv("SECRET_KEY", ""),
	}
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func InitConfig() error {
	return godotenv.Load()
}

var Settings Config
