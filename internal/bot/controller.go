package bot

import (
	"Kee-Reall/dnd-manager/internal/config"
	"Kee-Reall/dnd-manager/internal/domain"
	"Kee-Reall/dnd-manager/internal/service"
	"Kee-Reall/dnd-manager/internal/shared"
	"errors"
	"fmt"
	"log"
	"slices"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Context struct {
	update *tgbotapi.Update
	sender *tgbotapi.User
	user   *domain.User
}

type MethodSignature[T func(Context) tgbotapi.MessageConfig | func(Context, []string) tgbotapi.MessageConfig] struct {
	resolver T
	rights   []domain.Role
}

func (ms *MethodSignature[T]) Cred(role domain.Role) bool {
	return !(role > domain.MaxRole()) && slices.Contains(ms.rights, role)
}

type IdProvider interface {
	IdInInt64() (int64, bool)
	IdString() (string, bool)
}

type Controller struct {
	container *service.Container
	bot       *tgbotapi.BotAPI
	CmdList   map[string]MethodSignature[func(Context) tgbotapi.MessageConfig]
	CbList    map[string]MethodSignature[func(Context, []string) tgbotapi.MessageConfig]
}

func NewController(provider *service.Container, bot *tgbotapi.BotAPI) *Controller {
	c := &Controller{container: provider, bot: bot}

	c.CmdList = make(map[string]MethodSignature[func(Context) tgbotapi.MessageConfig])

	c.CmdList["start"] = MethodSignature[func(Context) tgbotapi.MessageConfig]{
		c.StartCMD,
		domain.AllRole(),
	}

	c.CmdList["master"] = MethodSignature[func(Context) tgbotapi.MessageConfig]{
		c.MasterKeyboardInit,
		domain.MasterFunctionsRole(),
	}

	c.CbList = make(map[string]MethodSignature[func(Context, []string) tgbotapi.MessageConfig])
	c.CbList["accept-reg"] = MethodSignature[func(Context, []string) tgbotapi.MessageConfig]{
		c.AcceptReg,
		[]domain.Role{domain.AdminRole},
	}

	//c.CbList["master-campaign-create"] = c.CreateCampaign

	c.container.EventBus.Subscribe("new-user-reg", func(payload any) error {
		m, ok := payload.(domain.User)
		if !ok {
			return errors.New("Can't cast payload to user")
		}
		go c.notifyAdminReg(m)
		return nil
	})

	c.container.EventBus.Subscribe("user-accepted", func(payload any) error {
		m, ok := payload.(domain.User)
		if !ok {
			return errors.New("Can't cast payload to user")
		}
		go c.notifyUserRegAccepted(m)
		return nil
	})

	return c
}

func (c *Controller) AcceptReg(ctx Context, args []string) tgbotapi.MessageConfig {
	adminChatId, ok := c.container.IVariable(config.AdminChatId)
	if args == nil || len(args) != 1 || !ok || adminChatId != int(ctx.sender.ID) {
		return c.UnknownQuery(ctx)
	}
	id := args[0]
	err := c.container.UserService().AcceptUser(id)
	if err != nil {
		if errors.Is(err, shared.ScenarioAlreadyDoneException) {
			return tgbotapi.NewMessage(int64(adminChatId), "Пользователь уже подтверждён")
		}
		return tgbotapi.NewMessage(int64(adminChatId), "Что то пошло не так "+err.Error())
	}
	return tgbotapi.NewMessage(int64(adminChatId), "Пользователь подтверждён")
}

func (c *Controller) notifyUserRegAccepted(du domain.User) {
	chatId, ok := du.Marker.IdInInt64()
	if !ok {
		return
	}
	c.Reply(tgbotapi.NewMessage(chatId, "Ваша регистрация подтверждена"))
}

/*
func (c *Controller) CreateCampaign(ctx Context) tgbotapi.MessageConfig {

}

*/

func (c *Controller) MasterKeyboardInit(ctx Context) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(ctx.sender.ID, "МАСТЕР В ЗДАНИИ")
	msg.ReplyMarkup = masterStartMarkup()
	return msg
}

func (c *Controller) notifyAdminReg(du domain.User) {
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
	msg.ReplyMarkup = adminConfirmMarkup(du.ID)
	c.Reply(msg)
}

func (r *Controller) Reply(m tgbotapi.Chattable) {
	defer func() {
		e := recover()
		if e != nil {
			log.Println(e)
		}
	}()

	_, err := r.bot.Send(m)
	if err != nil {
		log.Println(err.Error())
	}
	return
}

func (c *Controller) ShouldPassAs(user tgbotapi.User) (bool, *domain.User) {
	marker := &domain.UserMarker{ID: strconv.Itoa(int(user.ID)), Tag: "tg"}
	s := c.container.UserService()
	return s.AccessAndRoleByMarker(marker)
}

func (c *Controller) IDK(ctx Context) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(ctx.sender.ID, "Не понимаю, что вы имеете ввиду")
}

func (c *Controller) UnknownQuery(ctx Context) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(ctx.sender.ID, "неизвестный запрос")
}

func (c *Controller) UnknownCMD(ctx Context) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(ctx.sender.ID, "Нераспознанная команда")
}

func (c *Controller) StartCMD(ctx Context) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(ctx.sender.ID, "круто")
	msg.ReplyMarkup = startMarkup()
	return msg
}

func (c *Controller) NotAllowed(ctx Context) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(ctx.sender.ID, "Вам запрещенно это действие")
}

func (c *Controller) Registry(ctx *Context) tgbotapi.MessageConfig {
	v, ok := c.container.BVariable(config.RegEnable)
	if !ok || !v {
		_, err := c.container.UserService().UserByMarker(&domain.UserMarker{strconv.Itoa(int(ctx.sender.ID)), "tg"})
		if errors.Is(err, shared.DoesNotExistsException) {
			return tgbotapi.NewMessage(ctx.sender.ID, "Это частный бот. доступ запрещён")
		}
		return tgbotapi.NewMessage(ctx.sender.ID, "Ожидайте расмотрения от администрации")
	}

	err := c.container.UserService().RegisterNewUserByTag(strconv.Itoa(int(ctx.sender.ID)), ctx.sender.UserName)
	if err != nil {
		if errors.Is(err, shared.NotAllowedException) {
			return tgbotapi.NewMessage(ctx.sender.ID, "Ожидайте расмотрения от администрации")
		}
		return tgbotapi.NewMessage(ctx.sender.ID, "Что то пошло не так, попробуйте позже")
	}
	return tgbotapi.NewMessage(ctx.sender.ID, "Ваша заявка принята к расмотрению")
}
