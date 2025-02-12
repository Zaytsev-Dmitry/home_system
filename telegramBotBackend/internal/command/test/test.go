package test

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"telegramCLient/internal/command"
	"telegramCLient/internal/components/echo"
	"telegramCLient/internal/storage"
)

type TestCommand struct {
	messageStorage    storage.Storage
	component         *echo.Echo
	ctx               context.Context
	bot               *bot.Bot
	callbackHandlerID string
	action            command.Action
}

func NewTestCommand(action command.Action, st storage.Storage, bot *bot.Bot, ctx context.Context) *TestCommand {
	c := &TestCommand{
		messageStorage: st,
		bot:            bot,
		ctx:            ctx,
		action:         action,
	}
	textMeta := echo.TextMeta{
		ConfirmText: "Тестовый конфирм текст",
		StartText:   "Тестовый start текст",
	}

	c.component = echo.NewEcho(bot, c.getQuestions(), c.proceedResult, c.setUserInput, textMeta, []echo.Option{})
	return c
}

func (c *TestCommand) RegisterHandler() {
	c.callbackHandlerID = c.bot.RegisterHandler(bot.HandlerTypeMessageText, c.GetName(), bot.MatchTypeExact, c.callback)
}

func (T *TestCommand) proceedResult(result echo.Result) {
	fmt.Print(result)
}

func (t *TestCommand) setUserInput(userInput bool, chatId int64) {
	t.action.Log(chatId, t.GetName(), userInput, true)
}

func (c *TestCommand) ProceedUserAnswer(ctx context.Context, b *bot.Bot, update *models.Update) {

}

func (t *TestCommand) GetName() string {
	return "/test"
}

func (c *TestCommand) getQuestions() []echo.CollectItem {
	return []echo.CollectItem{
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
func (t *TestCommand) ClearState(chatId int64) {
}
