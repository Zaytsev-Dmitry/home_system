package handler

import (
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/callbackquery"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/message"
	"telegramCLient/external"
	"telegramCLient/external/dto"
	"telegramCLient/internal/bot"
	"telegramCLient/util"
)

type RegisterUserCommandHandler struct {
	AuthServerClient external.AuthServerClient
}

func NewRegisterUserCommandHandler(authUrl string) *RegisterUserCommandHandler {
	return &RegisterUserCommandHandler{
		AuthServerClient: *external.NewAuthServerClient(authUrl),
	}
}

var tempUsers = make(map[int64]User)

type State uint

const (
	StateDefault State = iota
	StateConfirm
)

// User data and state
type User struct {
	State     State
	FirstName string
	LastName  string
	Login     string
	Email     string
}

func (handler *RegisterUserCommandHandler) Init(dispatcher *ext.Dispatcher) {
	dispatcher.AddHandler(handlers.NewMessage(message.Text, handler.register))
	dispatcher.AddHandler(handlers.NewCallback(callbackquery.Equal("register_callback_yes"), handler.registerCallbackYes))
	dispatcher.AddHandler(handlers.NewCallback(callbackquery.Equal("register_callback_no"), handler.registerCallbackNo))
}

func (handler *RegisterUserCommandHandler) GetName() string {
	return "RegisterCommandHandler"
}

func (handler *RegisterUserCommandHandler) register(b *gotgbot.Bot, ctx *ext.Context) error {
	userID := ctx.Message.From.Id
	user := User{
		FirstName: ctx.Message.From.FirstName,
		LastName:  ctx.Message.From.LastName,
		Login:     ctx.Message.From.Username,
		Email:     ctx.Message.Text,
	}

	b.SendMessage(
		ctx.Message.Chat.Id,
		fmt.Sprintf(""+
			"Супер теперь подтверди почту \n"+
			"\n"+
			"Итак: %s...Это точно верная почта?\n%s", user.Login, user.Email),
		&gotgbot.SendMessageOpts{
			ReplyMarkup: handler.registerCallBackKeyboard(),
		})
	tempUsers[userID] = user
	return nil
}

func (handler *RegisterUserCommandHandler) registerCallbackYes(b *gotgbot.Bot, ctx *ext.Context) error {
	cb := ctx.Update.CallbackQuery
	user := tempUsers[cb.From.Id]

	var TG dto.CreateAccountRequestAccountType = "TG"
	savedUser := handler.AuthServerClient.RegisterUser(
		dto.CreateAccountRequest{
			AccountType: &TG,
			Email:       &user.Email,
			FirstName:   &cb.From.FirstName,
			LastName:    &cb.From.LastName,
			Login:       &cb.From.Username,
			TelegramId:  &cb.From.Id,
		})
	b.SendMessage(
		cb.Message.GetChat().Id,
		fmt.Sprintf("Добро пожаловать %s. \nТеперь тебе доступны следующие менюшки", savedUser.Login),
		&gotgbot.SendMessageOpts{
			ReplyMarkup: util.CreateMenuKeyboard(),
		})
	delete(tempUsers, cb.From.Id)
	return nil
}

func (handler *RegisterUserCommandHandler) registerCallbackNo(b *gotgbot.Bot, ctx *ext.Context) error {
	cb := ctx.Update.CallbackQuery
	delete(tempUsers, cb.From.Id)
	_, err2 := b.SendMessage(cb.Message.GetChat().Id, "Ну окей...погнали дальше. Введи свою почту заново", nil)
	bot.Dispatcher.AddHandler(handlers.NewMessage(message.Text, handler.register))
	if err2 != nil {
		return fmt.Errorf("failed to registerCallbackNo: %w", err2)
	}
	return nil
}

func (handler *RegisterUserCommandHandler) registerCallBackKeyboard() gotgbot.InlineKeyboardMarkup {
	return gotgbot.InlineKeyboardMarkup{
		InlineKeyboard: [][]gotgbot.InlineKeyboardButton{{
			{Text: "Да", CallbackData: "register_callback_yes"},
			{Text: "Нет", CallbackData: "register_callback_no"},
		}},
	}
}
