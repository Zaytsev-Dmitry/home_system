package echo

import "github.com/go-telegram/bot/models"

const (
	CONFIRM_YES = "yes"
	CONFIRM_NO  = "no"
)

func (e *Echo) buildDefaultStartKeyboard() models.ReplyMarkup {
	return &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "Все понятно", CallbackData: e.prefix},
			},
		},
	}
}

func (e *Echo) buildDefaultConfirmKeyboard() models.ReplyMarkup {
	return &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "Да", CallbackData: e.prefix + CONFIRM_YES},
				{Text: "Нет", CallbackData: e.prefix + CONFIRM_NO},
			},
		},
	}
}
