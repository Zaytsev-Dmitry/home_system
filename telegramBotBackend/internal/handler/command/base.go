package command

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type BaseCommand interface {
	Init() []bot.Option
	StartCommand(ctx context.Context, b *bot.Bot, update *models.Update)
	ProceedMessage(ctx context.Context, b *bot.Bot, update *models.Update)
	GetName() string
	ClearStatus(update *models.Update)
	AddToDelete(msg int)
}
