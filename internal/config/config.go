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
	RegEnable   bool
}

func Load() (*Config, error) {

	if err := godotenv.Load(); err != nil {
		return nil, errors.New("error loading environment")
	}
	AdminChatId, err := strconv.Atoi(os.Getenv("ADMIN_CHAT_ID"))
	if err != nil {
		return nil, errors.New("invalid ENV format")
	}

	var regEnable bool
	switch os.Getenv("REG_ENABLE") {
	case "true":
		regEnable = true
	case "false":
		regEnable = false
	default:
		return nil, errors.New("invalid ENV format")
	}

	cfg := Config{
		BotToken:    os.Getenv("BOT_TOKEN"),
		DBPath:      os.Getenv("DB_PATH"),
		AdminChatId: int64(AdminChatId),
		RegEnable:   regEnable,
	}
	return &cfg, nil
}
