package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

func LoadConfig() Config {
	godotenv.Load()

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "1323"
	}

	return Config{
		Port: port,
	}
}