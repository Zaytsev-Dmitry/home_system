package command

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"telegramCLient/util"
)

type MenuCommandHandler struct {
}

func NewMenuCommandHandler() *MenuCommandHandler {
	return &MenuCommandHandler{}
}

func (c *MenuCommandHandler) Init() []bot.Option {
	return []bot.Option{
		bot.WithMessageTextHandler("/menu", bot.MatchTypeExact, c.callback),
		bot.WithCallbackQueryDataHandler("open_menu", bot.MatchTypeExact, c.callback),
	}
}

func (handler *MenuCommandHandler) callback(ctx context.Context, b *bot.Bot, update *models.Update) {
	chatId, msgId := util.GetChatAndMsgId(update)
	//клик по кнопке
	if update.CallbackQuery != nil {
		b.EditMessageText(
			ctx,
			&bot.EditMessageTextParams{
				ChatID:      chatId,
				MessageID:   msgId,
				Text:        "Выбери то что тебе интересно",
				ReplyMarkup: handler.buildMenuKeyboard(),
			})
	} else {
		//была вызвана команда
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:      chatId,
			Text:        fmt.Sprintf("Выбери то что тебе интересно"),
			ReplyMarkup: handler.buildMenuKeyboard(),
		})
		b.DeleteMessage(ctx, &bot.DeleteMessageParams{
			ChatID:    chatId,
			MessageID: msgId,
		})
	}
}

func (handler *MenuCommandHandler) buildMenuKeyboard() models.ReplyMarkup {
	kb := &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "Записки 📅", CallbackData: "open_notes"},
				{Text: "Профиль 🤖", CallbackData: "open_profile"},
			},
			{
				{Text: "Учет совместных расходов 💸", CallbackData: "/expense_accounting"},
			},
		},
	}
	return kb
}
