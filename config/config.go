package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Global config instance
var AppConfig *Config

// Config holds the configuration values.
type Config struct {
	Port    string
	DB_PATH string
}

// TODO: This should be a singleton
// NewConfig initializes a new Config instance
func InitConfig() {
	cfg := &Config{}
	err := cfg.Load()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	AppConfig = cfg
}

// Load reads the environment variables and assigns them to Config
func (c *Config) Load() error {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Warning: No .env file found, using system env vars")
	}

	c.Port = os.Getenv("PORT")
	c.DB_PATH = os.Getenv("DB_PATH")

	log.Println("Config loaded successfully:", *c)

	return nil
}
