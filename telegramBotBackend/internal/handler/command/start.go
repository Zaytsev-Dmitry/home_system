package command

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"telegramCLient/internal/dao"
	"telegramCLient/util"
)

type StartCommandHandler struct {
	dao dao.TelegramBotDao
}

var tempMessageSlice = make(map[int64]TempMessage)
var msgToDelete = make([]int, 0)

type TempMessage struct {
	FirstName string
	LastName  string
	Username  string
	Email     string
	State     State
}

type State uint

const (
	StateDefault           State = iota
	StateDrawHelloKeyboard State = iota
	StateAskEmail
	StateConfirm
)

func NewStartCommandHandler(d dao.TelegramBotDao) *StartCommandHandler {
	return &StartCommandHandler{dao: d}
}

func (h *StartCommandHandler) Init() []bot.Option {
	return []bot.Option{
		bot.WithMessageTextHandler("/start", bot.MatchTypeExact, h.StartCommand),
		bot.WithCallbackQueryDataHandler("start_callback", bot.MatchTypeExact, h.callback),
		bot.WithCallbackQueryDataHandler("register_callback_yes", bot.MatchTypeExact, h.callback),
		bot.WithCallbackQueryDataHandler("register_callback_no", bot.MatchTypeExact, h.callback),
	}
}

func (h *StartCommandHandler) StartCommand(ctx context.Context, b *bot.Bot, update *models.Update) {
	chatId, msgId := util.GetChatAndMsgId(update)
	h.dao.ActionRepo.Save(chatId, "START_COMMAND", "", msgId)
	h.callback(ctx, b, update)
}

func (h *StartCommandHandler) ProceedMessage(ctx context.Context, b *bot.Bot, update *models.Update) {
	h.callback(ctx, b, update)
}

func (h *StartCommandHandler) GetName() string {
	return "START_COMMAND"
}

func (h *StartCommandHandler) buildKeyboard() models.ReplyMarkup {
	kb := &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "Понятно", CallbackData: "start_callback"},
			},
		},
	}
	return kb
}

func (h *StartCommandHandler) callback(ctx context.Context, b *bot.Bot, update *models.Update) {
	message := util.GetChatMessage(update)
	user := tempMessageSlice[message.Chat.ID]
	var keyboard models.ReplyMarkup
	var text string
	var isCanEdit = true
	switch user.State {
	case StateDefault:
		user.State = StateDrawHelloKeyboard
		text = "Привет сначала мне нужна твоя почта"
		keyboard = h.buildKeyboard()
		isCanEdit = false
		b.DeleteMessage(ctx, &bot.DeleteMessageParams{ChatID: message.Chat.ID, MessageID: message.ID})
	case StateDrawHelloKeyboard:
		user.State = StateAskEmail
		text = "Введи свою почту"
	case StateAskEmail:
		msgToDelete = append(msgToDelete, message.ID)
		user.Email = message.Text
		user.FirstName = message.Chat.FirstName
		user.LastName = message.Chat.LastName
		user.Username = message.Chat.Username

		keyboard = h.confirmKeyboard()
		user.State = StateConfirm
		text = fmt.Sprintf("Супер %s теперь конфирми почту\n"+
			"Это она? "+"%s", user.Username, user.Email)
		isCanEdit = false
	case StateConfirm:
		if update.CallbackQuery.Data == "register_callback_no" {
			text = "Ну окэй поехали дальше. Введи почту"
			user.State = StateAskEmail
		} else {
			b.DeleteMessages(ctx, &bot.DeleteMessagesParams{ChatID: message.Chat.ID, MessageIDs: msgToDelete})
			//TODO зарегать пользака и очистить диалог
			h.dao.ActionRepo.Update(message.Chat.ID, "", "", message.ID)
		}
	default:
		panic("unknown state")
	}

	tempMessageSlice[message.Chat.ID] = user
	if isCanEdit {
		b.EditMessageText(ctx, &bot.EditMessageTextParams{
			MessageID:   message.ID,
			ChatID:      message.Chat.ID,
			Text:        text,
			ReplyMarkup: keyboard,
		})
	} else {
		sendMessage, _ := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:      message.Chat.ID,
			Text:        text,
			ReplyMarkup: keyboard,
		})
		msgToDelete = append(msgToDelete, sendMessage.ID)
	}
}

func (handler *StartCommandHandler) confirmKeyboard() models.ReplyMarkup {
	kb := &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "Да", CallbackData: "register_callback_yes"},
				{Text: "Нет", CallbackData: "register_callback_no"},
			},
		},
	}

	return kb
}
