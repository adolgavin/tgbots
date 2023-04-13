package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	TelegramAPIKey string
	ChatGPTAPIKey  string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		TelegramAPIKey: os.Getenv("TELEGRAM_API_KEY"),
		ChatGPTAPIKey:  os.Getenv("CHATGPT_API_KEY"),
	}
}
