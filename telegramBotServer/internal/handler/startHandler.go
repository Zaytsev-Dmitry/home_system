package handler

import (
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/callbackquery"
)

type StartCommandHandler struct {
	tempMessageCollection map[int64][]int64
}

func NewStartCommandHandler(tempMsgColl map[int64][]int64) *StartCommandHandler {
	return &StartCommandHandler{
		tempMessageCollection: tempMsgColl,
	}
}

func (handler *StartCommandHandler) Init(dispatcher *ext.Dispatcher) {
	dispatcher.AddHandler(handlers.NewCommand("start", handler.start))
	dispatcher.AddHandler(handlers.NewCallback(callbackquery.Equal("start_callback"), handler.startCB))
}

// TODO отловить ошибки
func (handler *StartCommandHandler) start(b *gotgbot.Bot, ctx *ext.Context) error {
	handler.tempMessageCollection[ctx.Message.GetChat().Id] = append([]int64{}, ctx.Message.MessageId)
	message, _ := b.SendMessage(
		ctx.Message.Chat.Id,
		"Привет!Прежде чем начать работу мне необходима твоя почта! Нажми кнопку \"Вперед\" чтобы продолжить",
		&gotgbot.SendMessageOpts{
			ReplyMarkup: gotgbot.InlineKeyboardMarkup{
				InlineKeyboard: [][]gotgbot.InlineKeyboardButton{{
					{Text: "Вперед", CallbackData: "start_callback"},
				}},
			},
		})
	handler.tempMessageCollection[message.Chat.Id] = append(handler.tempMessageCollection[message.Chat.Id], message.MessageId)

	return nil
}

func (handler *StartCommandHandler) startCB(b *gotgbot.Bot, ctx *ext.Context) error {
	cb := ctx.Update.CallbackQuery
	message, err2 := b.SendMessage(cb.Message.GetChat().Id, "Отправь мне свой email", nil)
	if err2 != nil {
		return fmt.Errorf("failed to startCB: %w", err2)
	}
	handler.tempMessageCollection[message.Chat.Id] = append(handler.tempMessageCollection[message.Chat.Id], cb.Message.GetMessageId())
	handler.tempMessageCollection[message.Chat.Id] = append(handler.tempMessageCollection[message.Chat.Id], message.MessageId)
	return nil
}
