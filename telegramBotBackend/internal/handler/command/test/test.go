package test

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"telegramCLient/internal/components/echo"
	"telegramCLient/internal/handler/command"
	"telegramCLient/internal/storage"
)

type TestCommand struct {
	messageStorage    storage.Storage
	component         *echo.Echo
	ctx               context.Context
	bot               *bot.Bot
	callbackHandlerID string
	action            command.UserAction
}

func NewTestCommand(action command.UserAction, st storage.Storage, bot *bot.Bot, ctx context.Context) *TestCommand {
	c := &TestCommand{
		messageStorage: st,
		bot:            bot,
		ctx:            ctx,
		action:         action,
	}
	c.component = echo.NewEcho(bot, c.getQuestions(), c.proceedResult, c.LogCommandAction, "Тестовый start текст", "Тестовый конфирм текст", "Комплит тест текст")
	return c
}

func (c *TestCommand) RegisterHandler() {
	c.callbackHandlerID = c.bot.RegisterHandler(bot.HandlerTypeMessageText, c.GetName(), bot.MatchTypeExact, c.callback)
}

func (T *TestCommand) proceedResult(result []echo.CollectItem) {
	fmt.Print(result)
}

func (c *TestCommand) ProceedUserAnswer(ctx context.Context, b *bot.Bot, update *models.Update) {

}

func (t *TestCommand) LogCommandAction(userId int64, status string) {
	t.action.LogCommand(userId, status, t.GetName())
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
