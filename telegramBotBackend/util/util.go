package util

import (
	"encoding/json"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/go-telegram/bot/models"
	"net/http"
)

func InlineKeyboardButton(text string, callbackData string) gotgbot.InlineKeyboardButton {
	return gotgbot.InlineKeyboardButton{
		Text:         text,
		CallbackData: callbackData,
	}
}

func InlineKeyboardRow(buttons ...gotgbot.InlineKeyboardButton) []gotgbot.InlineKeyboardButton {
	return buttons
}

func InlineKeyboard(rows ...[]gotgbot.InlineKeyboardButton) *gotgbot.InlineKeyboardMarkup {
	return &gotgbot.InlineKeyboardMarkup{
		InlineKeyboard: rows,
	}
}

func CreateMenuKeyboard() *gotgbot.InlineKeyboardMarkup {
	return InlineKeyboard(
		InlineKeyboardRow(
			InlineKeyboardButton("Записки", "notes"),
			InlineKeyboardButton("Профиль", "profile"),
		),
		InlineKeyboardRow(
			InlineKeyboardButton("Совместный учет расходов", "accounting_of_expenses"),
		),
	)
}

func GetChatAndMsgId(update *models.Update) (int64, int) {
	var chatId int64
	var msgId int
	if update.Message != nil {
		chatId = update.Message.Chat.ID
		msgId = update.Message.ID
	} else {
		chatId = update.CallbackQuery.Message.Message.Chat.ID
		msgId = update.CallbackQuery.Message.Message.ID
	}
	return chatId, msgId
}

func ParseResponseToStruct(respBody *http.Response, response any) any {
	defer respBody.Body.Close()
	decoder := json.NewDecoder(respBody.Body)
	decoder.Decode(response)
	return response
}
