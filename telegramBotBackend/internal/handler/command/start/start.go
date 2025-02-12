package start

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"telegramCLient/external"
	"telegramCLient/internal/components/echo"
	"telegramCLient/internal/dao"
	"telegramCLient/internal/handler/loader"
	"telegramCLient/internal/storage"
)

type StartCommand struct {
	dao               dao.TelegramBotDao
	authServerClient  *external.AuthServerClient
	component         *echo.Echo
	messageStorage    storage.Storage
	ctx               context.Context
	bot               *bot.Bot
	callbackHandlerID string
}

func NewStartCommand(st storage.Storage, bot *bot.Bot, ctx context.Context, d dao.TelegramBotDao, serverClient *external.AuthServerClient) *StartCommand {
	s := &StartCommand{
		dao:              d,
		authServerClient: serverClient,
		messageStorage:   st,
		bot:              bot,
		ctx:              ctx,
	}
	s.component = echo.NewEcho(bot, s.getQuestions(), s.proceedResult, loader.StartCommandText, loader.RegisterConfirmText)
	return s
}

func (s *StartCommand) RegisterHandler() {
	s.callbackHandlerID = s.bot.RegisterHandler(bot.HandlerTypeMessageText, s.GetName(), bot.MatchTypeExact, s.callback)
}

func (s *StartCommand) proceedUserAnswer(ctx context.Context, b *bot.Bot, update *models.Update) {
	s.component.ProceedUserAnswer(ctx, b, update)
}

func (s *StartCommand) proceedResult(result []echo.CollectItem) {
	fmt.Print(result)
}

func (s *StartCommand) ProceedUserAnswer(ctx context.Context, b *bot.Bot, update *models.Update) {
	s.component.ProceedUserAnswer(ctx, b, update)
}

func (s *StartCommand) GetName() string {
	return "/start"
}

func (s *StartCommand) getQuestions() []echo.CollectItem {
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
