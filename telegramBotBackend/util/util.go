package util

import (
	"encoding/json"
	"github.com/go-telegram/bot/models"
	"net/http"
)

func GetChatAndMsgId(update *models.Update) (int64, int) {
	var chatId int64
	var msgId int
	if update.Message != nil {
		chatId = update.Message.Chat.ID
		msgId = update.Message.ID
	} else {
		chatId = update.CallbackQuery.Message.Message.Chat.ID
		msgId = update.CallbackQuery.Message.Message.ID
	}
	return chatId, msgId
}

func GetChatMessage(update *models.Update) models.Message {
	var message models.Message
	if update.Message != nil {
		message = *update.Message
	} else {
		message = *update.CallbackQuery.Message.Message
	}
	return message
}

func ParseResponseToStruct(respBody *http.Response, response any) any {
	defer respBody.Body.Close()
	decoder := json.NewDecoder(respBody.Body)
	decoder.Decode(response)
	return response
}
