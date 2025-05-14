package start

import (
	"botCoreService/internal/app/components/echo"
	"botCoreService/internal/app/service"
	"botCoreService/internal/infrastructure/transport/tg/commands"
	"botCoreService/internal/infrastructure/transport/tg/commands/registry"
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type StartCommand struct {
	registry          registry.CommandRegistry
	component         echo.Echo
	ctx               context.Context
	bot               *bot.Bot
	callbackHandlerID string
}

func NewStartCommand(bot *bot.Bot, ctx context.Context, commandRegistry registry.CommandRegistry, ss service.SessionStorage) commands.BaseCommand {
	s := StartCommand{
		bot:      bot,
		ctx:      ctx,
		registry: commandRegistry,
	}
	textMeta := echo.TextMeta{
		ConfirmText: "А теперь подтверди данные. Все ли корректно?",
		StartText:   "Начнем регаться!",
	}

	questions := []echo.QuestionItem{
		{
			Content:   "Как мне к тебе обращаться?",
			FieldDesc: "Обращение",
		},
		{
			Content:   "Введи свой Email",
			FieldDesc: "email",
		},
	}
	s.component = echo.NewEcho(bot, questions, textMeta, s.registry, s.proceedResult, ss)
	return &s
}

func (s *StartCommand) proceedResult(result echo.EchoResult) {
	//TODO вызов keycloak
	fmt.Println(result)
}

func (s *StartCommand) ProceedUserInput(ctx context.Context, b *bot.Bot, update *models.Update) {
	//TODO добавить логер
	s.component.ProceedUserInput(ctx, b, update)
}

func (s *StartCommand) RegisterHandler() {
	s.callbackHandlerID = s.bot.RegisterHandler(bot.HandlerTypeMessageText, s.GetName(), bot.MatchTypeExact, s.callback)
}

func (s *StartCommand) GetName() string {
	return "/start"
}
