package handler

import (
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/callbackquery"
)

type StartCommandHandler struct {
}

func (handler *StartCommandHandler) Init(dispatcher *ext.Dispatcher) {
	dispatcher.AddHandler(handlers.NewCommand("start", start))
	dispatcher.AddHandler(handlers.NewCallback(callbackquery.Equal("start_callback"), startCB))
}

func (handler *StartCommandHandler) GetName() string {
	return "StartCommandHandler"
}

func start(b *gotgbot.Bot, ctx *ext.Context) error {
	b.SendMessage(
		ctx.Message.Chat.Id,
		"Супер теперь надо зарегаться",
		&gotgbot.SendMessageOpts{
			ReplyMarkup: gotgbot.InlineKeyboardMarkup{
				InlineKeyboard: [][]gotgbot.InlineKeyboardButton{{
					{Text: "Вперед", CallbackData: "start_callback"},
				}},
			},
		})
	return nil
}

func startCB(b *gotgbot.Bot, ctx *ext.Context) error {
	cb := ctx.Update.CallbackQuery
	_, err2 := b.SendMessage(cb.Message.GetChat().Id, "Для работы мне нужен твой email. Напиши его мне", nil)
	if err2 != nil {
		return fmt.Errorf("failed to startCB: %w", err2)
	}
	return nil
}
