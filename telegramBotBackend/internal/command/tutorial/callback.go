package tutorial

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"telegramCLient/util"
)

func (t *TutorialCommand) callback(ctx context.Context, b *bot.Bot, update *models.Update) {
	message := util.GetChatMessage(update)
	t.action.Log(message.Chat.ID, t.GetName(), false, true)
	t.component.CreateAndRun(dialogNodes, ctx, b, update)
}
