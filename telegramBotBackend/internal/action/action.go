package action

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"telegramCLient/internal/dao"
	"telegramCLient/internal/handler/command"
	"telegramCLient/internal/handler/loader"
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

func (a *Action) getCommandHandlerByName(name string) command.BaseCommand {
	return a.commands[name]
}

func (a *Action) Proceed(ctx context.Context, b *bot.Bot, update *models.Update) {
	userId, _ := util.GetChatAndMsgId(update)
	actionEntity := a.dao.ActionRepo.GetByTgId(userId)
	lastRunCommandHandler := a.getCommandHandlerByName(actionEntity.CommandName)
	if !actionEntity.NeedUserAction {
		text := fmt.Sprintf(
			loader.UnnecessaryActionInfo,
			actionEntity.CommandName,
		)
		message, _ := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:    userId,
			Text:      text,
			ParseMode: models.ParseModeHTML,
		})
		lastRunCommandHandler.ClearStatus(update)
		lastRunCommandHandler.AddToDelete(message.ID)
	} else {
		lastRunCommandHandler.ProceedMessage(ctx, b, update)
	}
}
