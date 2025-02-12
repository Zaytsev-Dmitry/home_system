package questionnaire

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"telegramCLient/util"
)

func (s *Questionnaire) callback(ctx context.Context, b *bot.Bot, update *models.Update) {
	chatId, _ := util.GetChatAndMsgId(update)
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      chatId,
		Text:        "Тестовый текс",
		ReplyMarkup: s.buildDefaultStartKeyboard,
	})
}
