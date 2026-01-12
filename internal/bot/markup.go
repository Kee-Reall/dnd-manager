package bot

import tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func PlayerMarkup() tg.InlineKeyboardMarkup {
	return tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData("Нажми", "/comand"),
		),
	)
}
