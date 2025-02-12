package start

import (
	"context"
	"github.com/go-telegram/bot"
	"telegramCLient/external"
	"telegramCLient/internal/components/echo"
	"telegramCLient/internal/dao"
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
	return &StartCommand{
		dao:              d,
		authServerClient: serverClient,
		messageStorage:   st,
		bot:              bot,
		ctx:              ctx,
	}
}

func (h *StartCommand) RegisterHandler() {
	h.callbackHandlerID = h.bot.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, h.callback)
}
