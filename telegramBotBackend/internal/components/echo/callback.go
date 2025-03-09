package echo

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"strconv"
	"strings"
	"telegramCLient/internal/storage"
	"telegramCLient/util"
)

type Status struct {
	actualState      State
	questionIterator int
}

type State uint

var ActualStatus = map[int64]Status{}

const (
	DEFAULT State = iota
	ASK_FIELDS
	CONFIRM
)

func (e *Echo) callback(ctx context.Context, b *bot.Bot, update *models.Update) {
	message := util.GetChatMessage(update)
	messageType := storage.BOT
	if update.CallbackQuery == nil {
		messageType = storage.USER
	}
	e.addToStorage(message.Chat.ID, &message, messageType)
	status := ActualStatus[message.Chat.ID]
	var text string
	var keyboard models.ReplyMarkup
	var isComplete = false
	switch status.actualState {
	case DEFAULT:
		text = "Первый вопрос: " + e.question[status.questionIterator].Content
		status = Status{actualState: ASK_FIELDS}
		e.setUserInput(true, message.Chat.ID)
	case ASK_FIELDS:
		{
			e.question[status.questionIterator].Answer = message.Text
			status = Status{actualState: ASK_FIELDS, questionIterator: status.questionIterator + 1}

			if status.questionIterator == len(e.question) {
				//получили все ответы - меняем статус
				status = Status{actualState: CONFIRM, questionIterator: 0}
				text = fillConfirmText(e.text.ConfirmText, e.question)
				keyboard = e.buildDefaultConfirmKeyboard()
			} else {
				text = "Следующий вопрос: " + e.question[status.questionIterator].Content
			}
		}
	case CONFIRM:
		{
			e.setUserInput(false, message.Chat.ID)
			cmd := strings.TrimPrefix(update.CallbackQuery.Data, e.prefix)
			if cmd == CONFIRM_YES {
				isComplete = true
			} else {
				text = "Ну хорошо давай заново: " + e.question[0].Content
				status = Status{actualState: ASK_FIELDS}
			}
		}
	}
	ActualStatus[message.Chat.ID] = status
	if !isComplete {
		sendMessage, _ := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:      message.Chat.ID,
			Text:        text,
			ReplyMarkup: keyboard,
			ParseMode:   models.ParseModeHTML,
		})
		e.addToStorage(sendMessage.Chat.ID, sendMessage, messageType)
	} else {
		e.sendResult(isComplete, message)
	}
}

func fillConfirmText(confirmText string, question []CollectItem) string {
	var result = confirmText
	for i, value := range question {
		result = strings.Replace(result, "name_"+strconv.Itoa(i), value.FieldName, -1)
		result = strings.Replace(result, "value_"+strconv.Itoa(i), value.Answer, -1)
	}
	return result
}

func (e *Echo) sendResult(isComplete bool, message models.Message) {
	if isComplete {
		all := e.messageStorage.GetAll(message.Chat.ID)
		ids := make([]int, len(all))
		for i, value := range all {
			ids[i] = value.Id
		}
		e.proceedResult(Result{
			ChatId:        message.Chat.ID,
			UserFirstName: message.Chat.FirstName,
			UserLastname:  message.Chat.LastName,
			Username:      message.Chat.Username,
			Question:      e.question,
			MessagesIds:   ids,
		})
	}
}
