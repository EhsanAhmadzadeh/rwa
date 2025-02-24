package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var PORT string = GetEnv("PORT", "8080")
var DATABASE_URL string = GetEnv("DATABASE_URL", "wadb.db")

// LoadConfig loads environment variables from .env file
func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found")
	}
}

// GetEnv gets environment variables with a fallback
func GetEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
