package command

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

var commandsMap map[string]BaseCommand = map[string]BaseCommand{}

type UserAction struct {
}

func NewUserAction() *UserAction {
	return &UserAction{}
}

func (ua *UserAction) AddCommand(comm BaseCommand) {
	commandsMap[comm.GetName()] = comm
}

func (ua *UserAction) Proceed(ctx context.Context, bot *bot.Bot, update *models.Update) {
	//chatId, msgId := util.GetChatAndMsgId(update)
	commandsMap["/start"].ProceedUserAnswer(ctx, bot, update)
}

func (ua *UserAction) proceedUserAnswer() {

}
