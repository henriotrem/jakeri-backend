package utils

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func LoadDotEnv() {
	if strings.EqualFold(os.Getenv("API_MODE"), "release") {
		return
	}

	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}
