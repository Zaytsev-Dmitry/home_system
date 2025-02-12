package start

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"telegramCLient/util"
)

func (s *StartCommand) callback(ctx context.Context, b *bot.Bot, update *models.Update) {
	message := util.GetChatMessage(update)
	s.action.Log(message.Chat.ID, s.GetName(), false, true)
	s.component.Collect(ctx, b, update)
}
