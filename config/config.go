package config

import (
	"os"

	"github.com/joho/godotenv"
)

func Get(k string) string {
	if err := godotenv.Load(".env"); err != nil {
		panic("Error loading .env file")
	}
	return os.Getenv(k)
}
