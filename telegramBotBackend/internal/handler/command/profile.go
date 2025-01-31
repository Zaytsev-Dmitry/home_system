package command

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"strconv"
	"telegramCLient/external"
	"telegramCLient/internal/handler/loader"
)

type ProfileCommandHandler struct {
	AuthServerClient *external.AuthServerClient
}

func NewProfileCommandHandler(authServerClient *external.AuthServerClient) *ProfileCommandHandler {
	return &ProfileCommandHandler{
		AuthServerClient: authServerClient,
	}
}

func (h *ProfileCommandHandler) Init() []bot.Option {
	return []bot.Option{
		bot.WithMessageTextHandler("/profile", bot.MatchTypeExact, h.callback),
		bot.WithCallbackQueryDataHandler("open_profile", bot.MatchTypeExact, h.callback),
		bot.WithCallbackQueryDataHandler("close_profile", bot.MatchTypeExact, h.close),
	}
}

func (h *ProfileCommandHandler) close(ctx context.Context, b *bot.Bot, update *models.Update) {
	var messageId int
	var chatId int
	if update.Message != nil {
		messageId = update.Message.ID
		chatId = int(update.Message.Chat.ID)
	} else {
		messageId = update.CallbackQuery.Message.Message.ID
		chatId = int(update.CallbackQuery.Message.Message.Chat.ID)
	}
	b.DeleteMessage(ctx, &bot.DeleteMessageParams{
		ChatID:    chatId,
		MessageID: messageId,
	})
}

func (h *ProfileCommandHandler) callback(ctx context.Context, b *bot.Bot, update *models.Update) {
	var tgId int
	if update.Message != nil {
		tgId = int(update.Message.From.ID)
	} else {
		tgId = int(update.CallbackQuery.From.ID)
	}
	profile := h.AuthServerClient.GetProfileByTelegramId(tgId)

	text := fmt.Sprintf(
		loader.ProfileInfoDescText,
		*profile.TelegramUsername,
		strconv.Itoa(tgId),
		"Тут должен быть твой email, но его украли",
		*profile.Role,
	)
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      tgId,
		Text:        text,
		ReplyMarkup: h.buildKeyboard(),
		ParseMode:   models.ParseModeMarkdown,
	})
}

func (h *ProfileCommandHandler) buildKeyboard() models.ReplyMarkup {
	kb := &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "Закрыть ❌", CallbackData: "close_profile"},
				{Text: "Вызов меню 🤙", CallbackData: "open_menu"},
			},
		},
	}
	return kb
}
