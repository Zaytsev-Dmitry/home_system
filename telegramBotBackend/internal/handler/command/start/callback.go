package start

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"telegramCLient/util"
)

func (s *StartCommand) callback(ctx context.Context, b *bot.Bot, update *models.Update) {
	util.GetChatAndMsgId(update)
	s.component.Collect(ctx, b, update)
}
