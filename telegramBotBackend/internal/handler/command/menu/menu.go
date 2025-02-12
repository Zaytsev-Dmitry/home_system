package menu

import (
	"context"
	"github.com/go-telegram/bot"
	"telegramCLient/internal/storage"
)

type MenuCommand struct {
	messageStorage    storage.Storage
	ctx               context.Context
	bot               *bot.Bot
	callbackHandlerID string
}

func NewMenuCommand(st storage.Storage, bot *bot.Bot, ctx context.Context) *MenuCommand {
	return &MenuCommand{
		messageStorage: st,
		ctx:            ctx,
		bot:            bot,
	}
}

func (m *MenuCommand) RegisterHandler() {
	m.callbackHandlerID = m.bot.RegisterHandler(bot.HandlerTypeMessageText, "/menu", bot.MatchTypeExact, m.callback)
}
