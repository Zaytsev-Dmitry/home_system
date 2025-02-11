package paginator

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/go-telegram/ui/paginator"
)

type Paginator struct {
}

func NewPaginator() *Paginator {
	return &Paginator{}
}

func (*Paginator) CreateAndRun(ctx context.Context, b *bot.Bot, update *models.Update, data []string, perPage int, closeBtnText string) {
	opts := []paginator.Option{
		paginator.PerPage(perPage),
		paginator.WithoutEmptyButtons(),
		paginator.WithCloseButton(closeBtnText),
	}
	p := paginator.New(b, data, opts...)
	p.Show(ctx, b, update.CallbackQuery.Message.Message.Chat.ID)
}
