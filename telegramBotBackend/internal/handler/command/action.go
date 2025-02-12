package command

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"telegramCLient/internal/dao"
	"telegramCLient/internal/handler/loader"
	"telegramCLient/util"
)

type ActionItem struct {
	LastUserAction        string
	LastRequirementToUser string
	LastUserSentMessageId int
}
type Action struct {
	commands map[string]BaseCommand
	dao      dao.TelegramBotDao
}

func NewAction(dao dao.TelegramBotDao) *Action {
	return &Action{
		dao:      dao,
		commands: map[string]BaseCommand{},
	}
}

func (a *Action) AddCommand(c BaseCommand) {
	a.commands[c.GetName()] = c
}

func (a *Action) Log(tgId int64, cName string, userInput bool, isRunning bool) {
	a.dao.ActionRepo.SaveOrUpdate(tgId, userInput, 0, cName, isRunning)
}

func (a *Action) getCommandHandlerByName(name string) BaseCommand {
	return a.commands[name]
}

func (a *Action) Proceed(ctx context.Context, b *bot.Bot, update *models.Update) {
	userId, _ := util.GetChatAndMsgId(update)
	actionEntity := a.dao.ActionRepo.GetByTgId(userId)
	lastRunCommandHandler := a.getCommandHandlerByName(actionEntity.CommandName)
	if !actionEntity.NeedUserInput {
		text := fmt.Sprintf(
			loader.UnnecessaryActionText,
			actionEntity.CommandName,
		)
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:    userId,
			Text:      text,
			ParseMode: models.ParseModeHTML,
		})
		a.Log(userId, lastRunCommandHandler.GetName(), false, false)
		lastRunCommandHandler.ClearState(userId)
		//TODO заюзать message storage для update
	} else {
		lastRunCommandHandler.ProceedUserAnswer(ctx, b, update)
	}
}
