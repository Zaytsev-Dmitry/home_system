package test

import (
	"context"
	"github.com/go-telegram/bot"
	"telegramCLient/internal/components/questionnaire"
	"telegramCLient/internal/storage"
)

type TestCommand struct {
	messageStorage    storage.Storage
	component         *questionnaire.Questionnaire
	ctx               context.Context
	bot               *bot.Bot
	callbackHandlerID string
}

func NewTestCommand(st storage.Storage, bot *bot.Bot, ctx context.Context) *TestCommand {
	c := &TestCommand{
		messageStorage: st,
		bot:            bot,
		ctx:            ctx,
	}
	c.component = questionnaire.New(bot, c.getQuestions(), st, c.proceedResult)
	return c
}

func (c *TestCommand) RegisterHandler() {
	c.callbackHandlerID = c.bot.RegisterHandler(bot.HandlerTypeMessageText, "/test", bot.MatchTypeExact, c.callback)
}

func (c *TestCommand) proceedResult(result string) {

}

func (c *TestCommand) getQuestions() []questionnaire.CollectItem {
	return []questionnaire.CollectItem{
		{
			FieldId:   "username",
			FieldName: "Логин: ",
			Content:   "Как мне к тебе обращаться?",
		},
		{
			FieldId:   "email",
			FieldName: "Почта: ",
			Content:   "Введи свой Email",
		},
	}
}
