package start

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (s *StartCommand) callback(ctx context.Context, b *bot.Bot, update *models.Update) {
	//TODO команда запущена
	s.component.Collect(ctx, b, update)
}
