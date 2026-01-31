package bot

import (
	"Kee-Reall/dnd-manager/internal/domain"
	"Kee-Reall/dnd-manager/internal/service"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func NewRunner(bot *tgbotapi.BotAPI, container *service.Container) (*Runner, error) {
	r := &Runner{NewController(container, bot)}
	return r, nil
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

		ctx := Context{}
		ctx.update = &u

		if u.Message != nil {
			ctx.sender = u.Message.From
		} else {
			ctx.sender = u.CallbackQuery.From
		}

		pass, role := r.controller.ShouldPassAs(*ctx.sender)
		ctx.role = role
		if !pass {
			go r.controller.Reply(r.controller.Registry(&ctx))
			continue
		}

		if role == domain.NoRole {
			go r.controller.Reply(tgbotapi.NewMessage(ctx.sender.ID, "ожидайте подтверждения"))
			continue
		}

		if u.Message != nil { // хендлим команды
			go r.HandleMessage(ctx)
			continue
		}

		if u.CallbackQuery != nil { //  кнопки
			go r.HandleQuery(ctx)
			continue
		}
	}
}

func parseCallbackQuery(callbackData string) (string, []string) {
	if callbackData == "" {
		return "", make([]string, 0, 0)
	}
	split := strings.Split(callbackData, "_")
	for i, v := range split {
		split[i] = strings.TrimSpace(v)
	}
	return split[0], split[1:]
}

func (r *Runner) HandleQuery(ctx Context) {
	cmd, args := parseCallbackQuery(ctx.update.CallbackQuery.Data)
	resolver, ok := r.controller.CbList[cmd]
	if !ok {
		go r.controller.Reply(r.controller.UnknownQuery(*ctx.update))
		return
	}
	go r.controller.Reply(resolver(ctx, args))
}

func (r *Runner) HandleMessage(ctx Context) {
	if !ctx.update.Message.IsCommand() {
		go r.controller.Reply(r.controller.IDK(*ctx.update))
		return
	}

	cmd := ctx.update.Message.Command()
	resolver, ok := r.controller.CmdList[cmd]
	if !ok {
		go r.controller.Reply(r.controller.UnknownCMD(*ctx.update))
		return
	}
	go r.controller.Reply(resolver(ctx))
}
