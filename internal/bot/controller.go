package bot

import (
	"Kee-Reall/dnd-manager/internal/domain"
	"Kee-Reall/dnd-manager/internal/service"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

//пока заготовка, под более точную типизацию конкретных сервисов
/*type ServiceProvider interface {

}*/

type Controller struct {
	provider *service.Container
	bot      *tgbotapi.BotAPI
	CmdList  map[string]func(tgbotapi.Update) tgbotapi.MessageConfig
	CbList   map[string]func(tgbotapi.Update) tgbotapi.MessageConfig
}

func NewController(provider *service.Container, bot *tgbotapi.BotAPI) *Controller {
	c := &Controller{provider: provider, bot: bot}
	c.CmdList = make(map[string]func(tgbotapi.Update) tgbotapi.MessageConfig)
	c.CmdList["start"] = c.StartCMD
	return c
}

func (c *Controller) ShouldPassAs(user tgbotapi.User) (bool, domain.Role) {
	marker := domain.UserMarker{ID: strconv.Itoa(int(user.ID)), Tag: "tg"}
	s := c.provider.UserService()
	return s.AccessAndRoleByMarker(marker)
}

func (c *Controller) IDK(u tgbotapi.Update) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(u.Message.Chat.ID, "Не понимаю, что вы имеете ввиду")
}

func (c *Controller) UnknownQuery(u tgbotapi.Update) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(u.CallbackQuery.From.ID, "неизвестный запрос")
}

func (c *Controller) UnknownCMD(u tgbotapi.Update) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(u.Message.Chat.ID, "Нераспознанная команда")
}

func (c *Controller) StartCMD(u tgbotapi.Update) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(u.Message.Chat.ID, "круто")
	msg.ReplyMarkup = PlayerMarkup()
	return msg
}
