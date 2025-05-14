package telegram

import "github.com/go-telegram/bot/models"

func GetUserId(update *models.Update) int64 {
	if update.CallbackQuery != nil {
		return update.CallbackQuery.From.ID
	} else {
		return update.Message.From.ID
	}
}
