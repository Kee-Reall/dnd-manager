package bot

import (
	"Kee-Reall/dnd-manager/internal/config"
	"Kee-Reall/dnd-manager/internal/domain"
	"Kee-Reall/dnd-manager/internal/service"
	"errors"
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Context struct {
	update *tgbotapi.Update
	sender *tgbotapi.User
	role   domain.Role
}

type IdProvider interface {
	IdInInt64() (int64, bool)
	IdString() (string, bool)
}

type Controller struct {
	container *service.Container
	bot       *tgbotapi.BotAPI
	CmdList   map[string]func(Context) tgbotapi.MessageConfig
	CbList    map[string]func(Context, []string) tgbotapi.MessageConfig
}

func NewController(provider *service.Container, bot *tgbotapi.BotAPI) *Controller {
	c := &Controller{container: provider, bot: bot}

	c.CmdList = make(map[string]func(ctx Context) tgbotapi.MessageConfig)
	c.CmdList["start"] = c.StartCMD

	c.CbList = make(map[string]func(Context, []string) tgbotapi.MessageConfig)
	c.CbList["accept-reg"] = c.AcceptReg

	c.container.EventBus.Subscribe("new-user-reg", func(payload any, event string) error {
		m, ok := payload.(domain.User)
		if !ok {
			return errors.New("Can't cast payload to userMarker")
		}
		go c.NotifyAdminReg(m)
		return nil
	})

	return c
}

func (c *Controller) AcceptReg(ctx Context, args []string) tgbotapi.MessageConfig {
	adminChatId, ok := c.container.IVariable(config.AdminChatId)
	if !ok {
		return nil
	}
	ctx.sender.ID !== c.container.IVariable(config.AdminChatId)
	if ctx.role != domain.AdminRole {

	}
}

func (c *Controller) NotifyAdminReg(du domain.User) {
	chatId, ok := c.container.IVariable(config.AdminChatId)
	if !ok {
		return
	}
	uIdI64, ok := du.Marker.IdInInt64()
	if !ok {
		return
	}
	msg := tgbotapi.NewMessage(
		int64(chatId),
		fmt.Sprintf("<b>Новая заявка</b>\n <b>от пользователя %s</b>\n<a href='tg://user?id=%d'>профиль</a>", du.Name, uIdI64),
	)
	msg.ParseMode = "HTML"
	msg.ReplyMarkup = AdminConfirmMarkup(du.ID)
	c.Reply(msg)
}

func (r *Controller) Reply(m tgbotapi.Chattable) {
	_, err := r.bot.Send(m)
	defer func(e error) {
		if e != nil {
			log.Println(e.Error())
		}
	}(err)
	return
}

func (c *Controller) ShouldPassAs(user tgbotapi.User) (bool, domain.Role) {
	marker := domain.UserMarker{ID: strconv.Itoa(int(user.ID)), Tag: "tg"}
	s := c.container.UserService()
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

func (c *Controller) StartCMD(ctx Context) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(ctx.sender.ID, "круто")
	msg.ReplyMarkup = StartMarkup()
	return msg
}

func (c *Controller) Registry(ctx *Context) tgbotapi.MessageConfig {
	v, ok := c.container.BVariable(config.RegEnable)
	if !ok || !v {
		_, err := c.container.UserService().UserByMarker(&domain.UserMarker{strconv.Itoa(int(ctx.sender.ID)), "tg"})
		if errors.Is(err, domain.DoesNotExistsException) {
			return tgbotapi.NewMessage(ctx.sender.ID, "Это частный бот. доступ запрещён")
		}
		return tgbotapi.NewMessage(ctx.sender.ID, "Ожидайте расмотрения от администрации")
	}

	err := c.container.UserService().RegisterNewUserByTag(strconv.Itoa(int(ctx.sender.ID)), ctx.sender.UserName)
	if err != nil {
		if errors.Is(err, domain.NotAllowedException) {
			return tgbotapi.NewMessage(ctx.sender.ID, "Ожидайте расмотрения от администрации")
		}
		return tgbotapi.NewMessage(ctx.sender.ID, "Что то пошло не так, попробуйте позже")
	}
	return tgbotapi.NewMessage(ctx.sender.ID, "Ваша заявка принята к расмотрению")
}
