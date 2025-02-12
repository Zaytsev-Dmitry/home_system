package test

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (comm *TestCommand) callback(ctx context.Context, b *bot.Bot, update *models.Update) {
	fmt.Print("run TestCommand callback")
	comm.component.Collect()
}
