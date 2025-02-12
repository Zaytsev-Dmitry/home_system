package menu

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"telegramCLient/internal/handler/command"
	"telegramCLient/internal/storage"
)

type MenuCommand struct {
	messageStorage    storage.Storage
	ctx               context.Context
	bot               *bot.Bot
	callbackHandlerID string
	action            command.UserAction
}

func NewMenuCommand(action command.UserAction, st storage.Storage, bot *bot.Bot, ctx context.Context) *MenuCommand {
	return &MenuCommand{
		messageStorage: st,
		ctx:            ctx,
		bot:            bot,
		action:         action,
	}
}

func (m *MenuCommand) RegisterHandler() {
	m.callbackHandlerID = m.bot.RegisterHandler(bot.HandlerTypeMessageText, m.GetName(), bot.MatchTypeExact, m.callback)
}

func (p *MenuCommand) LogCommandAction(userId int64, status string) {
	p.action.LogCommand(userId, status, p.GetName())
}

func (m *MenuCommand) ProceedUserAnswer(ctx context.Context, b *bot.Bot, update *models.Update) {
}

func (m *MenuCommand) GetName() string {
	return "/menu"
}
