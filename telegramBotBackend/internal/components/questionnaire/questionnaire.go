package questionnaire

import (
	"fmt"
	"github.com/go-telegram/bot"
	"telegramCLient/internal/storage"
)

type proceedResult func(result string)

type Questionnaire struct {
	question          []CollectItem
	messageStorage    storage.Storage
	prefix            string
	callbackHandlerID string
	proceedResult     proceedResult
}

func New(b *bot.Bot, questions []CollectItem, storage storage.Storage, pr proceedResult) *Questionnaire {
	q := &Questionnaire{
		question:       questions,
		messageStorage: storage,
		prefix:         bot.RandomString(16),
		proceedResult:  pr,
	}
	q.callbackHandlerID = b.RegisterHandler(bot.HandlerTypeCallbackQueryData, q.prefix, bot.MatchTypePrefix, q.callback)
	return q
}

func (q *Questionnaire) Collect() {
	fmt.Print("run component collect")
}
