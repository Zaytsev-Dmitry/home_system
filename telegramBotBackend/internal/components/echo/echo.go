package echo

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"telegramCLient/internal/storage"
	"telegramCLient/util"
)

type proceedResult func(result Result)
type setUserInput func(userInput bool, chatId int64)

type Echo struct {
	question          []CollectItem
	messageStorage    storage.Storage
	prefix            string
	callbackHandlerID string
	proceedResult     proceedResult
	setUserInput      setUserInput
	text              TextMeta
}

type TextMeta struct {
	ConfirmText string
	StartText   string
}

func NewEcho(b *bot.Bot, questions []CollectItem, pr proceedResult, ui setUserInput, textP TextMeta, opts []Option) *Echo {
	e := &Echo{
		question:      questions,
		prefix:        bot.RandomString(16),
		proceedResult: pr,
		setUserInput:  ui,
		text:          textP,
	}
	for _, opt := range opts {
		opt(e)
	}

	e.callbackHandlerID = b.RegisterHandler(bot.HandlerTypeCallbackQueryData, e.prefix, bot.MatchTypePrefix, e.callback)
	return e
}

// TODO отловить ошибки
func (e *Echo) Collect(ctx context.Context, b *bot.Bot, update *models.Update) {
	sourceMessage := util.GetChatMessage(update)
	message, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      sourceMessage.Chat.ID,
		Text:        e.text.StartText,
		ReplyMarkup: e.buildDefaultStartKeyboard(),
		ParseMode:   models.ParseModeHTML,
	})
	if err != nil {
		fmt.Print(err.Error())
	}
	e.addToStorage(sourceMessage.Chat.ID, &sourceMessage, storage.USER)
	e.addToStorage(sourceMessage.Chat.ID, message, storage.BOT)
}

func (e *Echo) ProceedUserAnswer(ctx context.Context, b *bot.Bot, update *models.Update) {
	e.callback(ctx, b, update)
}

func (e *Echo) addToStorage(chatId int64, message *models.Message, messageType storage.MsgType) {
	if &e.messageStorage != nil {
		m := *storage.NewMessage(
			message.ID,
			message.Text,
			messageType,
		)
		e.messageStorage.Add(chatId, m)
	}
}

func (e *Echo) ClearState(chatId int64) {
	e.question = []CollectItem{}
	delete(ActualStatus, chatId)
}
