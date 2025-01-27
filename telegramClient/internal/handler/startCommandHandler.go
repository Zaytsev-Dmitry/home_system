package handler

import (
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

func HandleStartCommand(bh *th.BotHandler) {
	bh.HandleMessage(func(bot *telego.Bot, message telego.Message) {
		_, _ = bot.SendMessage(tu.Messagef(
			tu.ID(message.Chat.ID),
			"Привет %s!", message.From.Username+" "+"сначала тебе необходимо зарегистрироваться",
		).WithReplyMarkup(tu.InlineKeyboard(
			tu.InlineKeyboardRow(tu.InlineKeyboardButton("Go!").WithCallbackData("register"))),
		))
	}, th.CommandEqual("start"))

	bh.HandleCallbackQuery(func(bot *telego.Bot, query telego.CallbackQuery) {
		_, _ = bot.SendMessage(tu.Message(tu.ID(query.Message.GetChat().ID), "GO"))
		_ = bot.AnswerCallbackQuery(tu.CallbackQuery(query.ID).WithText("Done"))
	}, th.AnyCallbackQueryWithMessage(), th.CallbackDataEqual("register"))
}
