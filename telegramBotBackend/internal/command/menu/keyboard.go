package menu

import "github.com/go-telegram/bot/models"

func (c *MenuCommand) buildMenuKeyboard() models.ReplyMarkup {
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
