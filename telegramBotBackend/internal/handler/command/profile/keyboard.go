package profile

import "github.com/go-telegram/bot/models"

func (h *ProfileCommand) buildKeyboard() models.ReplyMarkup {
	return &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "Закрыть ❌", CallbackData: "close_profile"},
				{Text: "Вызов меню 🤙", CallbackData: "open_menu"},
			},
		},
	}
}
