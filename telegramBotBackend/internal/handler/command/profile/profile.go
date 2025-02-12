package profile

import (
	"context"
	"github.com/go-telegram/bot"
	"telegramCLient/external"
	"telegramCLient/internal/storage"
)

type ProfileCommand struct {
	AuthServerClient  *external.AuthServerClient
	messageStorage    storage.Storage
	ctx               context.Context
	bot               *bot.Bot
	callbackHandlerID string
}

func NewProfileCommand(st storage.Storage, bot *bot.Bot, ctx context.Context, authServerClient *external.AuthServerClient) *ProfileCommand {
	return &ProfileCommand{
		AuthServerClient: authServerClient,
		messageStorage:   st,
		bot:              bot,
		ctx:              ctx,
	}
}

func (p *ProfileCommand) RegisterHandler() {
	p.callbackHandlerID = p.bot.RegisterHandler(bot.HandlerTypeMessageText, "/profile", bot.MatchTypeExact, p.callback)
}
