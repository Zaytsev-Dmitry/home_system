package echo

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"telegramCLient/internal/dao/repository/intefraces"
)

type Echo struct {
	bot                     *bot.Bot
	ctx                     context.Context
	chatId                  int64
	firstSentMsgId          int
	prefix                  string
	callbackHandlerIDs      []string
	startKeyboard           models.ReplyMarkup
	confirmKeyboard         models.ReplyMarkup
	startKeyboardText       string
	confirmKeyboardText     string
	completeText            string
	questions               []CollectItem
	confirmCallbackFunction func(result Result)
	actionRepo              intefraces.ActionRepository
	commandName             string
}

var messagesToDelete []int
var tempDataSlice = make(map[int64]dataCollect)

func NewEcho(
	ctx context.Context,
	b *bot.Bot,
	chatId int64,
	startMsgId int,
	opts []Option,
	actionRepo intefraces.ActionRepository,
	commandName string,
) *Echo {
	p := &Echo{
		bot:         b,
		ctx:         ctx,
		chatId:      chatId,
		prefix:      bot.RandomString(16),
		actionRepo:  actionRepo,
		commandName: commandName,
	}
	messagesToDelete = append(messagesToDelete, startMsgId)
	for _, opt := range opts {
		opt(p)
	}
	return p
}

func (echo *Echo) StartCollect() {
	data := tempDataSlice[echo.chatId]
	data.State = StateDrawStartKeyboard
	message, _ := echo.bot.SendMessage(echo.ctx, &bot.SendMessageParams{
		ChatID:      echo.chatId,
		Text:        echo.startKeyboardText,
		ReplyMarkup: echo.startKeyboard,
		ParseMode:   models.ParseModeHTML,
	})
	echo.firstSentMsgId = message.ID
}

func (echo *Echo) ProceedAnswer(ctx context.Context, b *bot.Bot, update *models.Update) {
	echo.callback(ctx, b, update)
}
