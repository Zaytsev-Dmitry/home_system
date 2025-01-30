package components

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/go-telegram/ui/paginator"
	"strconv"
)

type Paginator struct {
}

func NewPaginator() *Paginator {
	return &Paginator{}
}

func (*Paginator) CreateAndRun(ctx context.Context, b *bot.Bot, update *models.Update, data []string, perPage int) {
	opts := []paginator.Option{
		paginator.PerPage(perPage),
		paginator.WithoutEmptyButtons(),
		paginator.WithCloseButton("Close"),
	}
	p := paginator.New(b, data, opts...)
	p.Show(ctx, b, strconv.Itoa(int(update.Message.Chat.ID)))
}
