package util

import "github.com/PaulSonOfLars/gotgbot/v2"

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
			InlineKeyboardButton("Напоминалка", "reminders"),
		),
		InlineKeyboardRow(
			InlineKeyboardButton("Совместный учет расходов", "accounting_of_expenses"),
		),
	)
}
