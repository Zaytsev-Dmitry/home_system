package echo

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"strings"
	"telegramCLient/util"
)

type State uint

const (
	StateDrawStartKeyboard State = iota
	StateAskFields
	StateConfirm
)

var answerIteratorIndex int = 0
var confirmCallbackYes = "callback_yes"
var confirmCallbackNo = "callback_no"
var startKeyboardCallback = "start_callback"

func (echo *Echo) callback(ctx context.Context, b *bot.Bot, update *models.Update) {
	message := util.GetChatMessage(update)
	data := tempDataSlice[message.Chat.ID]

	var keyboard models.ReplyMarkup
	var text string
	var isDoneCollect = false
	switch data.State {
	case StateDrawStartKeyboard:
		data.State = StateAskFields
		text = "Итак начнем..." + echo.getQuestion()
		echo.updateUserAction(true, "StateAskFields", message)
	case StateAskFields:
		//добавляем сообщения от пользователя чтобы потом их удалить
		messagesToDelete = append(messagesToDelete, message.ID)
		text = echo.collectAnswer(&data, message, text)
		if len(data.answers) == len(echo.questions) {
			keyboard, text = echo.collectAnswersIsDone(keyboard, &data)
			echo.updateUserAction(false, "StateConfirm", message)
		}
	case StateConfirm:
		//пользователь подтвердил введеные данные
		if strings.TrimPrefix(update.CallbackQuery.Data, echo.prefix) == confirmCallbackYes {
			//шлем результат в вызвавщий компонент
			text = echo.proceedConfirmYes(message, data, text, b)
			isDoneCollect = true
			echo.updateUserAction(false, "done", message)
		} else {
			text = echo.confirmProceedNo(&data, text)
			isDoneCollect = true
			echo.updateUserAction(true, "StateAskFields", message)
		}
	}

	tempDataSlice[message.Chat.ID] = data
	if isDoneCollect {
		//удалить ранее отправленные
		b.DeleteMessages(echo.ctx, &bot.DeleteMessagesParams{
			ChatID:     echo.chatId,
			MessageIDs: messagesToDelete,
		})
		//редактирую самое первое сообщение
		b.EditMessageText(echo.ctx, &bot.EditMessageTextParams{
			ChatID:    echo.chatId,
			MessageID: echo.firstSentMsgId,
			Text:      text,
			ParseMode: models.ParseModeHTML,
		})
	} else {
		sendMessage, _ := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:      message.Chat.ID,
			Text:        text,
			ReplyMarkup: keyboard,
			ParseMode:   models.ParseModeHTML,
		})
		messagesToDelete = append(messagesToDelete, sendMessage.ID)
	}
}

func (echo *Echo) collectAnswer(data *dataCollect, message models.Message, text string) string {
	//продолжаем собирать ответы
	echo.addAnswer(data, message)
	text = echo.getQuestion()
	return text
}

func (echo *Echo) collectAnswersIsDone(keyboard models.ReplyMarkup, data *dataCollect) (models.ReplyMarkup, string) {
	//собрали все ответы
	keyboard = echo.confirmKeyboard
	keyboardText := echo.confirmKeyboardText
	for i, answer := range data.answers {
		id := answer.FieldId
		content := data.answers[i].Content
		keyboardText = strings.Replace(keyboardText, id, content, -1)
	}
	data.State = StateConfirm
	return keyboard, keyboardText
}

func (echo *Echo) confirmProceedNo(data *dataCollect, text string) string {
	//ретрай опрашивания данных
	data.State = StateAskFields
	data.answers = []CollectItem{}
	echo.restoreIterator()
	text = "Ну окей...давай заново: " + echo.getQuestion()
	return text
}

func (echo *Echo) proceedConfirmYes(message models.Message, data dataCollect, text string, b *bot.Bot) string {
	echo.confirmCallbackFunction(Result{
		ChatId:        echo.chatId,
		MsgId:         message.ID,
		Answers:       data.answers,
		UserFirstName: message.Chat.FirstName,
		UserLastname:  message.Chat.LastName,
		UserTGName:    message.Chat.Username,
	})
	text = echo.completeText
	for _, uid := range echo.callbackHandlerIDs {
		b.UnregisterHandler(uid)
	}
	return text
}

func (echo *Echo) restoreIterator() {
	answerIteratorIndex = 0
}

func (echo *Echo) getQuestion() string {
	return echo.questions[answerIteratorIndex].Content
}

func (echo *Echo) updateUserAction(needUserAction bool, state string, message models.Message) {
	echo.actionRepo.SaveOrUpdate(message.Chat.ID, state, needUserAction, message.ID, echo.commandName)
}

func (echo *Echo) addAnswer(data *dataCollect, message models.Message) {
	data.answers = append(data.answers, CollectItem{
		FieldId:   echo.questions[answerIteratorIndex].FieldId,
		FieldName: echo.questions[answerIteratorIndex].FieldName,
		Content:   message.Text,
	})
	if answerIteratorIndex+1 != len(echo.questions) {
		answerIteratorIndex = answerIteratorIndex + 1
	}
}
