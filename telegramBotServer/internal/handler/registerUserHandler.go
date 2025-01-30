package handler

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"telegramCLient/external"
	"telegramCLient/external/dto"
)

type RegisterUserCommandHandler struct {
	AuthServerClient      external.AuthServerClient
	tempMessageCollection map[int64][]int
}

func NewRegisterUserCommandHandler(authUrl string, tempMsgColl map[int64][]int) *RegisterUserCommandHandler {
	return &RegisterUserCommandHandler{
		AuthServerClient:      *external.NewAuthServerClient(authUrl),
		tempMessageCollection: tempMsgColl,
	}
}

var tempUsers = make(map[int64]User)

type User struct {
	FirstName string
	LastName  string
	Login     string
	Email     string
}

func (h *RegisterUserCommandHandler) Init() []bot.Option {
	return []bot.Option{
		bot.WithDefaultHandler(h.register),
		bot.WithCallbackQueryDataHandler("register_callback_yes", bot.MatchTypeExact, h.registerCallbackYes),
		bot.WithCallbackQueryDataHandler("register_callback_no", bot.MatchTypeExact, h.registerCallbackNo),
		bot.WithDefaultHandler(h.register),
	}
}

// TODO отловаить ошибки
func (handler *RegisterUserCommandHandler) register(ctx context.Context, b *bot.Bot, update *models.Update) {
	var message = update.Message
	if message == nil {
		message = update.CallbackQuery.Message.Message
	}

	handler.tempMessageCollection[message.From.ID] = append(handler.tempMessageCollection[message.From.ID], message.ID)
	userID := message.From.ID
	user := User{
		FirstName: message.Chat.FirstName,
		LastName:  message.Chat.LastName,
		Login:     message.Chat.Username,
		Email:     message.Text,
	}

	sendedMSG, _ := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: message.Chat.ID,
		Text: fmt.Sprintf(""+
			"Супер теперь подтверди почту \n"+
			"\n"+
			"Итак: %s...Это точно верная почта?\n%s", user.Login, user.Email),
		ReplyMarkup: handler.buildKeyboard(),
	})
	tempUsers[userID] = user
	handler.tempMessageCollection[message.From.ID] = append(handler.tempMessageCollection[message.From.ID], sendedMSG.ID)
}

// TODO отловаить ошибки
func (handler *RegisterUserCommandHandler) registerCallbackYes(ctx context.Context, b *bot.Bot, update *models.Update) {
	cb := update.CallbackQuery
	user := tempUsers[cb.From.ID]
	var TG dto.CreateAccountRequestAccountType = "TG"
	savedUser := handler.AuthServerClient.RegisterUser(
		dto.CreateAccountRequest{
			AccountType: &TG,
			Email:       &user.Email,
			FirstName:   &cb.From.FirstName,
			LastName:    &cb.From.LastName,
			Login:       &cb.From.Username,
			TelegramId:  &cb.From.ID,
		})
	delete(tempUsers, cb.From.ID)
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: cb.Message.Message.Chat.ID,
		Text: fmt.Sprintf("Добро пожаловать %s. \nТеперь тебе доступны следующие менюшки. "+
			"И чтобы узнать о том что я умею. Просто нажми кнопку \"Помощь\" и следуй инструкциям", savedUser.Login),
		ReplyMarkup: handler.buildMenuKeyboard(),
	})

	b.DeleteMessages(ctx, &bot.DeleteMessagesParams{
		ChatID:     cb.Message.Message.Chat.ID,
		MessageIDs: handler.tempMessageCollection[cb.Message.Message.Chat.ID],
	})
	//очищаю мапу
	delete(handler.tempMessageCollection, cb.Message.Message.Chat.ID)
}

// TODO отловаить ошибки
func (handler *RegisterUserCommandHandler) registerCallbackNo(ctx context.Context, b *bot.Bot, update *models.Update) {
	cb := update.CallbackQuery
	delete(tempUsers, cb.From.ID)
	sendedMsg, _ := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: cb.Message.Message.Chat.ID,
		Text: fmt.Sprintf("" +
			"Ну окей...погнали дальше. Введи свою почту заново"),
	})
	handler.tempMessageCollection[cb.Message.Message.Chat.ID] = append(handler.tempMessageCollection[cb.Message.Message.Chat.ID], cb.Message.Message.ID)
	handler.tempMessageCollection[cb.Message.Message.Chat.ID] = append(handler.tempMessageCollection[cb.Message.Message.Chat.ID], sendedMsg.ID)
}

func (handler *RegisterUserCommandHandler) buildKeyboard() models.ReplyMarkup {
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

func (handler *RegisterUserCommandHandler) buildMenuKeyboard() models.ReplyMarkup {
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
