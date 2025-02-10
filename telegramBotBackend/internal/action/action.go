package action

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"telegramCLient/internal/handler/command"
	"telegramCLient/util"
)

var actionItems map[int64]*ActionItem

type ActionItem struct {
	LastUserAction        string
	LastRequirementToUser string
	LastUserSentMessageId int
}
type Action struct {
	commands map[string]command.BaseCommand
}

func NewAction() *Action {
	return &Action{}
}

func UpdateAction(userId int64, item ActionItem) {
	actionItems[userId] = &item
}

func (a *Action) Proceed(ctx context.Context, b *bot.Bot, update *models.Update) {
	userId, _ := util.GetChatAndMsgId(update)
	if actionItems[userId] != nil {
		lastAction := actionItems[userId].LastUserAction
		if lastAction == "START" {
			baseCommand := a.commands["START_COMMAND"]
			baseCommand.Proceed(ctx, b, update)
		}
	} else {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: userId,
			Text:   "Сори я не знаю что это такое",
		})
	}
}
