package config

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	BotToken string
	DBPath   string
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, errors.New("error loading environment")
	}
	cfg := Config{
		BotToken: os.Getenv("BOT_TOKEN"),
		DBPath:   os.Getenv("DB_PATH"),
	}
	return &cfg, nil
}
