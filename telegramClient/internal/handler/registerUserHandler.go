package handler

import (
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/callbackquery"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/message"
	"telegramCLient/internal/bot"
	"telegramCLient/internal/util"
)

type RegisterUserCommandHandler struct{}

var users = make(map[int64]User)

type State uint

const (
	StateDefault State = iota
	StateAskEmail
	StateConfirm
)

// User data and state
type User struct {
	State State
	Name  string
	Email string
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
	user := users[userID]

	var text string
	switch user.State {
	case StateDefault:
		text = "Введи свой email"
		user.Name = ctx.Message.Text
		user.State = StateAskEmail
	case StateAskEmail:
		user.Email = ctx.Message.Text
		user.State = StateConfirm
		users[userID] = user
		b.SendMessage(
			ctx.Message.Chat.Id,
			fmt.Sprintf("Login: %s \nEmail: %s \nВсе верно?", user.Name, user.Email),
			&gotgbot.SendMessageOpts{
				ReplyMarkup: handler.registerCallBackKeyboard(),
			})
		break
	default:
		panic("unknown state")
	}

	users[userID] = user
	_, err := b.SendMessage(ctx.Message.Chat.Id, text, nil)
	if err != nil {
		return fmt.Errorf("failed to echo message: %w", err)
	}
	return nil
}

// TODO вызов сервиса регистрации
func (handler *RegisterUserCommandHandler) registerCallbackYes(b *gotgbot.Bot, ctx *ext.Context) error {
	cb := ctx.Update.CallbackQuery
	user := users[cb.From.Id]
	b.SendMessage(
		cb.Message.GetChat().Id,
		fmt.Sprintf("Добро пожаловать %s. \nТеперь тебе доступны следующие менюшки", user.Name),
		&gotgbot.SendMessageOpts{
			ReplyMarkup: util.CreateMenuKeyboard(),
		})
	return nil
}

func (handler *RegisterUserCommandHandler) registerCallbackNo(b *gotgbot.Bot, ctx *ext.Context) error {
	cb := ctx.Update.CallbackQuery
	_, err2 := b.SendMessage(cb.Message.GetChat().Id, "Ну окей...погнали дальше. Введи свой логин", nil)
	delete(users, cb.From.Id)
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
