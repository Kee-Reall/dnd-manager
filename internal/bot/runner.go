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

		pass, user := r.controller.ShouldPassAs(*ctx.sender)
		ctx.user = user
		if !pass {
			resolve(func() {
				r.controller.Reply(r.controller.Registry(&ctx))
			})
			continue
		}

		if user.Role == domain.NoRole {
			resolve(func() {
				r.controller.Reply(tgbotapi.NewMessage(ctx.sender.ID, "ожидайте подтверждения"))
			})
			continue
		}

		if u.Message != nil { // хендлим команды
			resolve(func() {
				r.HandleMessage(ctx)
			})
			continue
		}

		if u.CallbackQuery != nil { //  кнопки
			resolve(func() {
				r.HandleQuery(ctx)
			})
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
	signature, ok := r.controller.CbList[cmd]
	if !ok {
		resolve(func() {
			r.controller.Reply(r.controller.UnknownQuery(ctx))
		})
		return
	}

	if !signature.Cred(ctx.user.Role) {
		resolve(func() {
			r.controller.Reply(r.controller.NotAllowed(ctx))
		})
		return
	}
	resolve(func() {
		r.controller.Reply(signature.resolver(ctx, args))
	})
}

func (r *Runner) HandleMessage(ctx Context) {

	if !ctx.update.Message.IsCommand() {
		resolve(func() {
			r.controller.Reply(r.controller.IDK(ctx))
		})
		return
	}

	cmd := ctx.update.Message.Command()
	signature, ok := r.controller.CmdList[cmd]
	if !ok {
		resolve(func() {
			r.controller.Reply(r.controller.UnknownCMD(ctx))
		})
		return
	}

	if !signature.Cred(ctx.user.Role) {
		resolve(func() {
			r.controller.Reply(r.controller.NotAllowed(ctx))
		})
		return
	}

	resolve(func() {
		r.controller.Reply(signature.resolver(ctx))
	})
}
