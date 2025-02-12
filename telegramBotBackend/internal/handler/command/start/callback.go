package start

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (s *StartCommand) callback(ctx context.Context, b *bot.Bot, update *models.Update) {
	s.component.Collect(ctx, b, update)
}
