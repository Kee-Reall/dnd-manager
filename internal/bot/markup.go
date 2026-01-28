package bot

import tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func StartMarkup() tg.InlineKeyboardMarkup {
	return tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData("Мои компании", "myCampaign"),
		),
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData("Ближайшие игры", "myGames"),
		),
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData("Открытые компании", "openCampaign"),
		),
	)
}
