package bot

import (
	"fmt"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartMarkup() tg.InlineKeyboardMarkup {
	return tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData("Мои компании", "my-campaign"),
		),
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData("Ближайшие игры", "my-games"),
		),
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData("Открытые компании", "open-campaign"),
		),
	)
}

func AdminConfirmMarkup(id string) tg.InlineKeyboardMarkup {
	return tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData("✅Подвердить", fmt.Sprintf("accept-reg_%s", id)),
			tg.NewInlineKeyboardButtonData("❌Отклонить", fmt.Sprintf("reject-reg_%s", id)),
		),
	)
}
