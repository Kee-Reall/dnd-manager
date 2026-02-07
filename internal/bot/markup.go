package bot

import (
	"fmt"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func startMarkup() tg.InlineKeyboardMarkup {
	return tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData("🌐Мои компании", "my-campaign"),
		),
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData("⏱️Ближайшие игры", "my-games"),
		),
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData("📋Открытые компании", "open-campaign"),
		),
	)
}

func adminConfirmMarkup(id string) tg.InlineKeyboardMarkup {
	return tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData("✅Подвердить", fmt.Sprintf("accept-reg_%s", id)),
			tg.NewInlineKeyboardButtonData("❌Отклонить", fmt.Sprintf("reject-reg_%s", id)),
		),
	)
}

func masterStartMarkup() tg.InlineKeyboardMarkup {
	return tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData("⚖️Заявки", "master-show-bid"),
		),
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData("🗓Список компаний", "master-campaign-list"),
			tg.NewInlineKeyboardButtonData("🔨Создать компанию", "master-campaign-create"),
		),
	)
}
