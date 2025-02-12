package echo

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"telegramCLient/internal/storage"
	"telegramCLient/util"
)

type proceedResult func(result []CollectItem)
type logCommand func(userId int64, status string)

type Echo struct {
	question          []CollectItem
	messageStorage    storage.Storage
	prefix            string
	callbackHandlerID string
	proceedResult     proceedResult
	logCommand        logCommand
	confirmText       string
	startText         string
	completeText      string
}

func NewEcho(b *bot.Bot, questions []CollectItem, pr proceedResult, lc logCommand, startText string, confirmText string, completeText string) *Echo {
	e := &Echo{
		question:      questions,
		prefix:        bot.RandomString(16),
		proceedResult: pr,
		logCommand:    lc,
		startText:     startText,
		confirmText:   confirmText,
		completeText:  completeText,
	}
	e.callbackHandlerID = b.RegisterHandler(bot.HandlerTypeCallbackQueryData, e.prefix, bot.MatchTypePrefix, e.callback)
	return e
}

// TODO отловить ошибки
func (e *Echo) Collect(ctx context.Context, b *bot.Bot, update *models.Update) {
	chatId, _ := util.GetChatAndMsgId(update)
	message, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      chatId,
		Text:        e.startText,
		ReplyMarkup: e.buildDefaultStartKeyboard(),
		ParseMode:   models.ParseModeHTML,
	})
	if err != nil {
		fmt.Print(err.Error())
	}
	e.addToStorage(chatId, message)
	e.logCommand(chatId, "start")
}

func (e *Echo) ProceedUserAnswer(ctx context.Context, b *bot.Bot, update *models.Update) {
	e.callback(ctx, b, update)
}

func (e *Echo) addToStorage(chatId int64, message *models.Message) {
	if &e.messageStorage != nil {
		m := *storage.NewMessage(
			message.ID,
			message.Text,
			0,
			storage.BOT,
		)
		e.messageStorage.Add(chatId, m)
	}
}
