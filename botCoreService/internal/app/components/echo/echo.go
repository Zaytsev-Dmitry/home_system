package echo

import (
	"botCoreService/internal/app/service"
	"botCoreService/internal/infrastructure/transport/tg/commands/registry"
	util "botCoreService/pkg"
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type Echo struct {
	CommandRegistry   registry.CommandRegistry
	Question          []QuestionItem
	prefix            string
	callbackHandlerID string
	text              TextMeta
	proceedResult     proceedResult
	sessionStorage    service.SessionStorage
}

// функция обработки результатов опроса
type proceedResult func(result EchoResult)

// TODO сделать поддержку кастомной клавиатуры (e.buildDefaultStartKeyboard())
func NewEcho(b *bot.Bot, questions []QuestionItem, textP TextMeta, registry registry.CommandRegistry, ps proceedResult, sessionStorage service.SessionStorage) Echo {
	e := Echo{
		CommandRegistry: registry,
		Question:        questions,
		prefix:          bot.RandomString(16),
		text:            textP,
		proceedResult:   ps,
		sessionStorage:  sessionStorage,
	}
	e.callbackHandlerID = b.RegisterHandler(bot.HandlerTypeCallbackQueryData, e.prefix, bot.MatchTypePrefix, e.callback)
	return e
}

func (e *Echo) Start(ctx context.Context, b *bot.Bot, update *models.Update) {
	sourceMessage := util.GetChatMessage(update)
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      sourceMessage.Chat.ID,
		Text:        e.text.StartText,
		ReplyMarkup: e.buildDefaultStartKeyboard(),
		ParseMode:   models.ParseModeHTML,
	})
	if err != nil {
		fmt.Print(err.Error())
	}
}

func (e *Echo) ProceedUserInput(ctx context.Context, b *bot.Bot, update *models.Update) {
	e.callback(ctx, b, update)
}
