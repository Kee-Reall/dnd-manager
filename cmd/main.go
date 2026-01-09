package main

import (
	"Kee-Reall/dnd-manager/internal/bot"
	"Kee-Reall/dnd-manager/internal/config"
	"Kee-Reall/dnd-manager/internal/service"
	"Kee-Reall/dnd-manager/internal/storage/sqlite"

	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	botAPI, err := tgbotapi.NewBotAPI(cfg.BotToken)
	if err != nil {
		log.Fatal(err)
	}

	db, err := sqlite.New(cfg.DBPath)
	if err != nil {
		log.Fatal(err)
	}

	services, err := service.NewContainer(db)
	if err != nil {
		log.Fatal(err)
	}

	if runner, err := bot.NewRunner(botAPI, services); err != nil {
		log.Fatal(err)
	} else {
		runner.Run()
	}
}
