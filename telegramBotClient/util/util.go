package util

import (
	"encoding/json"
	"github.com/PaulSonOfLars/gotgbot/v2"
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

func ParseResponseToStruct(respBody *http.Response, response any) any {
	defer respBody.Body.Close()
	decoder := json.NewDecoder(respBody.Body)
	decoder.Decode(response)
	return response
}
