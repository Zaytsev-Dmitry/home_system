package commands

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type BaseCommand interface {
	ProceedUserInput(ctx context.Context, b *bot.Bot, update *models.Update)
	RegisterHandler()
	GetName() string
}
