package start

import (
	"context"
	"errors"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"telegramCLient/external"
	"telegramCLient/util"
)

func (s *StartCommand) callback(ctx context.Context, b *bot.Bot, update *models.Update) {
	message := util.GetChatMessage(update)
	saved, err := s.authServerClient.GetAccountByTgId(message.Chat.ID)
	if err != nil && errors.Is(err, external.UNKNOWN) {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: message.Chat.ID,
			Text:   "Упс...произошла ошибка!Но админы уже знают об этом. Попробуй чуть позже",
		})
		return
	}
	if saved.ID != nil {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: message.Chat.ID,
			Text:   "Упс...оказыватся ты уже зареган. Соответственно тебе доступны все функции",
		})
	} else {
		s.action.Log(message.Chat.ID, s.GetName(), false, true)
		s.component.Collect(ctx, b, update)
	}
}
