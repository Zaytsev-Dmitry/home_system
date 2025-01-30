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

func (handler *StartCommandHandler) GetName() string {
	return "StartCommandHandler"
}

// TODO отловить ошибки
func (handler *StartCommandHandler) start(b *gotgbot.Bot, ctx *ext.Context) error {
	message, _ := b.SendMessage(
		ctx.Message.Chat.Id,
		"Супер теперь надо зарегаться",
		&gotgbot.SendMessageOpts{
			ReplyMarkup: gotgbot.InlineKeyboardMarkup{
				InlineKeyboard: [][]gotgbot.InlineKeyboardButton{{
					{Text: "Вперед", CallbackData: "start_callback"},
				}},
			},
		})
	tempMsgSlice := handler.tempMessageCollection[message.Chat.Id]
	if len(tempMsgSlice) == 0 {
		handler.tempMessageCollection[message.Chat.Id] = append([]int64{}, message.MessageId)
	} else {
		handler.tempMessageCollection[message.Chat.Id] = append(handler.tempMessageCollection[message.Chat.Id], message.MessageId)
	}

	return nil
}

func (handler *StartCommandHandler) startCB(b *gotgbot.Bot, ctx *ext.Context) error {
	cb := ctx.Update.CallbackQuery
	message, err2 := b.SendMessage(cb.Message.GetChat().Id, "Для работы мне нужен твой email. Напиши его мне", nil)
	if err2 != nil {
		return fmt.Errorf("failed to startCB: %w", err2)
	}
	handler.tempMessageCollection[message.Chat.Id] = append(handler.tempMessageCollection[message.Chat.Id], cb.Message.GetMessageId())
	handler.tempMessageCollection[message.Chat.Id] = append(handler.tempMessageCollection[message.Chat.Id], message.MessageId)
	return nil
}
