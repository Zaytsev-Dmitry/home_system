package notes

import (
	"context"
	"github.com/go-telegram/bot"
	"telegramCLient/external"
	"telegramCLient/internal/components/echo"
	"telegramCLient/internal/dao"
	"telegramCLient/internal/storage"
)

type NoteCommand struct {
	noteBackClient    *external.NoteBackendClient
	echoComponent     *echo.Echo
	dao               dao.TelegramBotDao
	messageStorage    storage.Storage
	ctx               context.Context
	bot               *bot.Bot
	callbackHandlerID string
}

func NewNotesCommand(st storage.Storage, bot *bot.Bot, ctx context.Context, d dao.TelegramBotDao, client *external.NoteBackendClient) *NoteCommand {
	return &NoteCommand{
		noteBackClient: client,
		dao:            d,
		messageStorage: st,
		ctx:            ctx,
		bot:            bot,
	}
}

func (n *NoteCommand) RegisterHandler() {
	n.callbackHandlerID = n.bot.RegisterHandler(bot.HandlerTypeMessageText, "/notes", bot.MatchTypeExact, n.callback)
}
