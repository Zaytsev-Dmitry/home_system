package command

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
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
	var chatId int64
	if update.Message != nil {
		chatId = update.Message.Chat.ID
	} else {
		chatId = update.CallbackQuery.Message.Message.Chat.ID
	}
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      chatId,
		Text:        fmt.Sprintf("Выбери то что тебе интересно"),
		ReplyMarkup: handler.buildMenuKeyboard(),
	})
}

func (handler *MenuCommandHandler) buildMenuKeyboard() models.ReplyMarkup {
	kb := &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "Записки 📅", CallbackData: "open_notes"},
				{Text: "Профиль 🤖", CallbackData: "/profile"},
			},
			{
				{Text: "Учет совместных расходов 💸", CallbackData: "/expense_accounting"},
			},
		},
	}
	return kb
}
