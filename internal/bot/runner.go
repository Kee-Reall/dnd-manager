package bot

import (
	"Kee-Reall/dnd-manager/internal/service"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func NewRunner(bot *tgbotapi.BotAPI, container *service.Container) (*Runner, error) {
	return &Runner{NewController(container, bot)}, nil
}

func postLog(e error) {
	if e != nil {
		log.Println(e.Error())
	}
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
	for u := range r.controller.bot.GetUpdatesChan(tgbotapi.NewUpdate(0)) {
		var from tgbotapi.User
		if u.Message != nil {
			from = *u.Message.From
		} else {
			from = *u.CallbackQuery.From
		}

		pass, role := r.controller.ShouldPassAs(from)

		if !pass {
			r.Reply(tgbotapi.NewMessage(from.ID, "Это частный бот, доступ запрещён"))
			continue
		}

		fmt.Printf("passing is %t, with role %s\n", pass, role.String())

		if u.Message != nil { // хендлим команды
			r.HandleMessage(u)
			continue
		}

		if u.CallbackQuery != nil { //  кнопки
			r.HandleQuery(u)
			continue
		}
	}
}

func (r *Runner) Reply(m tgbotapi.Chattable) {
	_, err := r.controller.bot.Send(m)
	defer postLog(err)
	return
}

func (r *Runner) HandleQuery(update tgbotapi.Update) {
	r.Reply(r.controller.UnknownQuery(update))
}

func (r *Runner) HandleMessage(update tgbotapi.Update) {
	if !update.Message.IsCommand() {
		r.Reply(r.controller.IDK(update))
		return
	}

	cmd := update.Message.Command()
	resolver, ok := r.controller.CmdList[cmd]
	if !ok {
		r.Reply(r.controller.UnknownCMD(update))
		return
	}
	r.Reply(resolver(update))
}
