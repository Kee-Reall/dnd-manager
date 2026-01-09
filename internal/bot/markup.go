package bot

import tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func PlayerMarkup() tg.ReplyKeyboardMarkup {
	return tg.NewReplyKeyboard([]tg.KeyboardButton{
		tg.NewKeyboardButton("Мои игры"),
	})
}
