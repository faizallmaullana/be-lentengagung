package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	App struct {
		Port string
	}
}

var config *Config

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found: %v", err)
	}

	config = &Config{}

	config.App.Port = getEnvWithDefault("APP_PORT", "9000")
}

func getEnvWithDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func Get() *Config {
	return config
}
