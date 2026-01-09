package bot

import (
	"Kee-Reall/dnd-manager/internal/service"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

//пока заготовка, под более точную типизацию конкретных сервисов
/*type ServiceProvider interface {

}*/

type Controller struct {
	provider *service.Container
	bot      *tgbotapi.BotAPI
	List     map[string]func(update tgbotapi.Update)
}

func setHandlersToController(c *Controller) {
	list := make(map[string]func(update tgbotapi.Update))
	list["start"] = c.StartCMD
	c.List = list
}

func NewController(provider *service.Container, bot *tgbotapi.BotAPI) *Controller {
	controller := &Controller{provider: provider, bot: bot}
	setHandlersToController(controller)
	return controller
}

func (c *Controller) UnknownCMD(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Нераспознанная команда")
	c.bot.Send(msg)
}

func (c *Controller) StartCMD(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "круто")
	msg.ReplyMarkup = PlayerMarkup()
	c.bot.Send(msg)
}
