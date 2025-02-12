package notes

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
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
	n.callbackHandlerID = n.bot.RegisterHandler(bot.HandlerTypeMessageText, n.GetName(), bot.MatchTypeExact, n.callback)
}

func (n *NoteCommand) ProceedUserAnswer(ctx context.Context, b *bot.Bot, update *models.Update) {
}

func (n *NoteCommand) GetName() string {
	return "/notes"
}
