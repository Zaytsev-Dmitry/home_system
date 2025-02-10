package action

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"telegramCLient/internal/dao"
	"telegramCLient/internal/handler/command"
	"telegramCLient/util"
)

type ActionItem struct {
	LastUserAction        string
	LastRequirementToUser string
	LastUserSentMessageId int
}
type Action struct {
	commands map[string]command.BaseCommand
	dao      dao.TelegramBotDao
}

func NewAction(dao dao.TelegramBotDao, commands []command.BaseCommand) *Action {
	a := &Action{dao: dao}
	comList := make(map[string]command.BaseCommand)
	for _, value := range commands {
		comList[value.GetName()] = value
	}
	a.commands = comList
	return a
}

func (a *Action) Proceed(ctx context.Context, b *bot.Bot, update *models.Update) {
	userId, _ := util.GetChatAndMsgId(update)
	actionEntity := a.dao.ActionRepo.GetByTgId(userId)
	if actionEntity.LastAction == "START_COMMAND" {
		baseCommand := a.commands["START_COMMAND"]
		baseCommand.ProceedMessage(ctx, b, update)
	} else {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: userId,
			Text:   "Сори я не знаю что это такое",
		})
	}
}
