package command

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type BaseCommand interface {
	RegisterHandler()
	ProceedUserAnswer(ctx context.Context, b *bot.Bot, update *models.Update)
	GetName() string
}
