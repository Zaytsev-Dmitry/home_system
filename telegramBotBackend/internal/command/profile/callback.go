package profile

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (p *ProfileCommand) callback(ctx context.Context, b *bot.Bot, update *models.Update) {
	//chatId, _ := util.GetChatAndMsgId(update)
	//p.action.Log(chatId, p.GetName(), false, true)
	//profile := p.AuthServerClient.GetProfileByTelegramId(int(chatId))
	//
	//var username string
	//if profile.TelegramUsername == nil {
	//	username = "Его украли но обещали скоро вернуть"
	//} else {
	//	username = *profile.TelegramUsername
	//}
	//text := fmt.Sprintf(
	//	loader.ProfileCommandContentText,
	//	username,
	//	chatId,
	//	"Тут должен быть твой email, но его украли",
	//	*profile.Role,
	//)
	//b.SendMessage(ctx, &bot.SendMessageParams{
	//	ChatID:      chatId,
	//	Text:        text,
	//	ReplyMarkup: p.buildKeyboard(),
	//	ParseMode:   models.ParseModeMarkdown,
	//})
}
