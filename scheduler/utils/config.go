package utils

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadConfigFromEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Print("Failed to load .env file: %v", err)
	}
}
