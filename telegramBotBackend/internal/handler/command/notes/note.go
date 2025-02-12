package notes

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"telegramCLient/external"
	"telegramCLient/internal/components/echo"
	"telegramCLient/internal/dao"
	"telegramCLient/internal/handler/command"
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
	action            command.UserAction
}

func NewNotesCommand(action command.UserAction, st storage.Storage, bot *bot.Bot, ctx context.Context, d dao.TelegramBotDao, client *external.NoteBackendClient) *NoteCommand {
	return &NoteCommand{
		noteBackClient: client,
		dao:            d,
		messageStorage: st,
		ctx:            ctx,
		bot:            bot,
		action:         action,
	}
}

func (n *NoteCommand) RegisterHandler() {
	n.callbackHandlerID = n.bot.RegisterHandler(bot.HandlerTypeMessageText, n.GetName(), bot.MatchTypeExact, n.callback)
}

func (p *NoteCommand) LogCommandAction(userId int64, status string) {
	p.action.LogCommand(userId, status, p.GetName())
}

func (n *NoteCommand) ProceedUserAnswer(ctx context.Context, b *bot.Bot, update *models.Update) {
}

func (n *NoteCommand) GetName() string {
	return "/notes"
}
