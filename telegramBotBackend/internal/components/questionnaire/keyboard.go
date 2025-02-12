package questionnaire

import "github.com/go-telegram/bot/models"

func (q *Questionnaire) buildDefaultStartKeyboard() models.ReplyMarkup {
	return &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "Все понятно", CallbackData: "next"},
			},
		},
	}
}

func (q *Questionnaire) buildDefaultConfirmKeyboard() models.ReplyMarkup {
	return &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "Да", CallbackData: "yes"},
				{Text: "Нет", CallbackData: "no"},
			},
		},
	}
}
