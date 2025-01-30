package components

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/go-telegram/ui/dialog"
)

type DialogInline struct {
}

func NewDialogInline() *DialogInline {
	return &DialogInline{}
}

func (*DialogInline) CreateAndRun(dialogNodes []dialog.Node, ctx context.Context, b *bot.Bot, update *models.Update) {
	p := dialog.New(b, dialogNodes, dialog.Inline())
	p.Show(ctx, b, update.Message.Chat.ID, "start")
}
