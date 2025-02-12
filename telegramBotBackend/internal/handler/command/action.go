package command

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"telegramCLient/util"
)

var commandsMap = map[string]BaseCommand{}
var actualUserCommand = map[int64]CommandState{}

type UserAction struct {
}

type CommandState struct {
	status      string
	commandName string
}

func NewUserAction() *UserAction {
	return &UserAction{}
}

func (ua *UserAction) AddCommand(comm BaseCommand) {
	commandsMap[comm.GetName()] = comm
}

func (ua *UserAction) LogCommand(userTgId int64, status string, commName string) {
	actualUserCommand[userTgId] = CommandState{
		status:      status,
		commandName: commName,
	}
	fmt.Print("dfgg")
}

func (ua *UserAction) Proceed(ctx context.Context, bot *bot.Bot, update *models.Update) {
	chatId, _ := util.GetChatAndMsgId(update)
	actualCommand := actualUserCommand[chatId]
	commandsMap[actualCommand.commandName].ProceedUserAnswer(ctx, bot, update)
}
