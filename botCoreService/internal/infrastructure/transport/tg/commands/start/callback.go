package start

import (
	"botCoreService/pkg/telegram"
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (s *StartCommand) callback(ctx context.Context, b *bot.Bot, update *models.Update) {
	s.registry.Set(telegram.GetUserId(update), s)
	s.component.Start(ctx, b, update)
}
