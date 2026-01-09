package bot

import (
	"Kee-Reall/dnd-manager/internal/service"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func NewRunner(bot *tgbotapi.BotAPI, container *service.Container) (*Runner, error) {
	return &Runner{NewController(container, bot)}, nil
}

type Runner struct {
	controller *Controller
}

func (r *Runner) Start() {
	go func() {
		r.Run()
	}()
}

func (r *Runner) Run() {
	for updating := range r.controller.bot.GetUpdatesChan(tgbotapi.NewUpdate(0)) {
		r.Handle(updating)
	}
}

func (r *Runner) Handle(update tgbotapi.Update) {
	cmd := update.Message.Command()
	resolver, ok := r.controller.List[cmd]
	if !ok {
		r.controller.UnknownCMD(update)
		return
	}
	resolver(update)
}
