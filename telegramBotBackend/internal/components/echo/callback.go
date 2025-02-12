package echo

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"strconv"
	"strings"
	"telegramCLient/util"
)

type Status struct {
	actualState      State
	questionIterator int
}

type State uint

var actualStatus = map[int64]Status{}

const (
	DEFAULT State = iota
	ASK_FIELDS
	CONFIRM
)

func (e *Echo) callback(ctx context.Context, b *bot.Bot, update *models.Update) {
	message := util.GetChatMessage(update)
	status := actualStatus[message.Chat.ID]
	var text string
	var keyboard models.ReplyMarkup
	switch status.actualState {
	case DEFAULT:
		text = "Первый вопрос: " + e.question[status.questionIterator].Content
		status = Status{actualState: ASK_FIELDS}
	case ASK_FIELDS:
		{
			e.question[status.questionIterator].Answer = message.Text
			status = Status{actualState: ASK_FIELDS, questionIterator: status.questionIterator + 1}

			if status.questionIterator == len(e.question) {
				//получили все ответы - меняем статус
				status = Status{actualState: CONFIRM, questionIterator: 0}

				for i, question := range e.question {
					e.confirmText = strings.Replace(e.confirmText, "name_"+strconv.Itoa(i), question.FieldName, -1)
					e.confirmText = strings.Replace(e.confirmText, "value_"+strconv.Itoa(i), question.Answer, -1)
				}
				text = e.confirmText
				keyboard = e.buildDefaultConfirmKeyboard()
			} else {
				text = "Следующий вопрос: " + e.question[status.questionIterator].Content
			}
		}
	case CONFIRM:
		{
			cmd := strings.TrimPrefix(update.CallbackQuery.Data, e.prefix)
			if cmd == CONFIRM_YES {
				e.proceedResult(e.question)
				e.logCommand(message.Chat.ID, "complete")
				text = e.completeText
			} else {
				text = "Ну хорошо давай заново: " + e.question[0].Content
				status = Status{actualState: ASK_FIELDS}
			}

		}
	}

	actualStatus[message.Chat.ID] = status

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      message.Chat.ID,
		Text:        text,
		ReplyMarkup: keyboard,
		ParseMode:   models.ParseModeHTML,
	})
}
