package handler

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type StartCommandHandler struct {
	tempMessageCollection map[int64][]int
}

func NewStartCommandHandler(tempMsgColl map[int64][]int) *StartCommandHandler {
	return &StartCommandHandler{
		tempMessageCollection: tempMsgColl,
	}
}

func (h *StartCommandHandler) Init() []bot.Option {
	return []bot.Option{
		bot.WithMessageTextHandler("/start", bot.MatchTypeExact, h.start),
		bot.WithCallbackQueryDataHandler("start_callback", bot.MatchTypeExact, h.startCB),
	}
}

func (h *StartCommandHandler) buildKeyboard() models.ReplyMarkup {
	kb := &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "Вперед:)", CallbackData: "start_callback"},
			},
		},
	}

	return kb
}

// TODO отловить ошибки
func (h *StartCommandHandler) start(ctx context.Context, b *bot.Bot, update *models.Update) {
	h.tempMessageCollection[update.Message.Chat.ID] = append([]int{}, update.Message.ID)

	message, _ := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Привет!Прежде чем начать работу мне необходима твоя почта! Нажми кнопку \"Вперед\" чтобы продолжить",
		ReplyMarkup: h.buildKeyboard(),
	})
	h.tempMessageCollection[update.Message.Chat.ID] = append(h.tempMessageCollection[update.Message.Chat.ID], message.ID)
}

// TODO отловить ошибки
func (h *StartCommandHandler) startCB(ctx context.Context, b *bot.Bot, update *models.Update) {
	h.tempMessageCollection[update.CallbackQuery.Message.Message.Chat.ID] = append(h.tempMessageCollection[update.CallbackQuery.Message.Message.Chat.ID], update.CallbackQuery.Message.Message.ID)
	sendedMSG, _ := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.CallbackQuery.Message.Message.Chat.ID,
		Text:   "Отправь мне свой email",
	})
	h.tempMessageCollection[update.CallbackQuery.Message.Message.Chat.ID] = append(h.tempMessageCollection[update.CallbackQuery.Message.Message.Chat.ID], sendedMSG.ID)
}
