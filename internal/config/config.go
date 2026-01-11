package config

import (
	"errors"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	BotToken    string
	DBPath      string
	AdminChatId int64
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, errors.New("error loading environment")
	}
	AdminChatId, err := strconv.Atoi(os.Getenv("ADMIN_CHAT_ID"))
	if err != nil {
		return nil, errors.New("invalid ENV format")
	}
	cfg := Config{os.Getenv("BOT_TOKEN"), os.Getenv("DB_PATH"), int64(AdminChatId)}
	return &cfg, nil
}
