package util

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	return godotenv.Load()
}

func GetAuthDomain() string {
	return os.Getenv("AUTH0_DOMAIN")
}

func GetAuthAudience() string {
	return os.Getenv("AUTH0_AUDIENCE")
}
