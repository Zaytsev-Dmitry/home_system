package profile

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"telegramCLient/external"
	"telegramCLient/internal/handler/command"
	"telegramCLient/internal/storage"
)

type ProfileCommand struct {
	AuthServerClient  *external.AuthServerClient
	messageStorage    storage.Storage
	ctx               context.Context
	bot               *bot.Bot
	callbackHandlerID string
	action            command.UserAction
}

func NewProfileCommand(action command.UserAction, st storage.Storage, bot *bot.Bot, ctx context.Context, authServerClient *external.AuthServerClient) *ProfileCommand {
	return &ProfileCommand{
		AuthServerClient: authServerClient,
		messageStorage:   st,
		bot:              bot,
		ctx:              ctx,
		action:           action,
	}
}

func (p *ProfileCommand) RegisterHandler() {
	p.callbackHandlerID = p.bot.RegisterHandler(bot.HandlerTypeMessageText, p.GetName(), bot.MatchTypeExact, p.callback)
}

func (p *ProfileCommand) LogCommandAction(userId int64, status string) {
	p.action.LogCommand(userId, status, p.GetName())
}

func (p *ProfileCommand) ProceedUserAnswer(ctx context.Context, b *bot.Bot, update *models.Update) {
}

func (p *ProfileCommand) GetName() string {
	return "/profile"
}
