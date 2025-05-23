package menu

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"telegramCLient/util"
)

func (comm *MenuCommand) callback(ctx context.Context, b *bot.Bot, update *models.Update) {
	chatId, msgId := util.GetChatAndMsgId(update)
	comm.action.Log(chatId, comm.GetName(), false, true)
	//клик по кнопке
	if update.CallbackQuery != nil {
		b.EditMessageText(
			ctx,
			&bot.EditMessageTextParams{
				ChatID:      chatId,
				MessageID:   msgId,
				Text:        "Выбери то что тебе интересно",
				ReplyMarkup: comm.buildMenuKeyboard(),
			})
	} else {
		//была вызвана команда
		message, _ := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:      chatId,
			Text:        fmt.Sprintf("Выбери то что тебе интересно"),
			ReplyMarkup: comm.buildMenuKeyboard(),
		})
		b.DeleteMessage(ctx, &bot.DeleteMessageParams{
			ChatID:    chatId,
			MessageID: msgId,
		})
		comm.action.Log(message.Chat.ID, comm.GetName(), false, true)
	}
}
