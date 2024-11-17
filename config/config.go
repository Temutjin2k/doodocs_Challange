package config

import (
	"log"

	"github.com/joho/godotenv"
)

var (
	AvailiableMimeTypesToSendEmail = map[string]bool{
		"application/vnd.openxmlformats - officedocument.wordprocessingml.document": true,
		"application/pdf": true,
	}

	AvailiableMimeTypesToArvhive = map[string]bool{
		"application/vnd.openxmlformats-officedocument.wordprocessingml.document": true,
		"application/xml": true,
		"image/jpeg":      true,
		"image/png":       true,
	}
)

func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}
