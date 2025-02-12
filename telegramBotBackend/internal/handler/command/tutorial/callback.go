package tutorial

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (t *TutorialCommand) callback(ctx context.Context, b *bot.Bot, update *models.Update) {
	t.component.CreateAndRun(dialogNodes, ctx, b, update)
}
