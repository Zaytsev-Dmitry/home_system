package test

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"telegramCLient/util"
)

func (t *TestCommand) callback(ctx context.Context, b *bot.Bot, update *models.Update) {
	message := util.GetChatMessage(update)
	t.action.Log(message.Chat.ID, t.GetName(), false, true)
	t.component.Collect(ctx, b, update)
}
