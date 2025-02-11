package command

import (
	"bytes"
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"telegramCLient/external"
	"telegramCLient/external/dto"
	"telegramCLient/internal/components/echo"
	"telegramCLient/internal/dao"
	"telegramCLient/internal/handler/loader"
	"telegramCLient/util"
)

type StartCommandHandler struct {
	dao              dao.TelegramBotDao
	authServerClient external.AuthServerClient
	echoComponent    echo.Echo
}

var tempMessageSlice = make(map[int64]TempUser)

type TempUser struct {
	FirstName string
	LastName  string
	Username  string
	Email     string
	State     State
}

type State uint

const (
	StateDefault State = iota
	StateDrawHelloKeyboard
	StateAskEmail
	StateConfirm
)

func NewStartCommandHandler(d dao.TelegramBotDao) *StartCommandHandler {
	return &StartCommandHandler{dao: d}
}

func (h *StartCommandHandler) Init() []bot.Option {
	return []bot.Option{
		bot.WithMessageTextHandler("/start", bot.MatchTypeExact, h.StartCommand),
		//bot.WithCallbackQueryDataHandler("start_callback", bot.MatchTypeExact, h.callback),
		//bot.WithCallbackQueryDataHandler("register_callback_yes", bot.MatchTypeExact, h.callback),
		//bot.WithCallbackQueryDataHandler("register_callback_no", bot.MatchTypeExact, h.callback),
	}
}

func (h *StartCommandHandler) StartCommand(ctx context.Context, b *bot.Bot, update *models.Update) {
	chatId, msgId := util.GetChatAndMsgId(update)
	////TODO проверить зареган ли пользак или нет
	//h.dao.ActionRepo.SaveOrUpdate(chatId, "StateDefault", false, msgId, h.GetName())
	//h.callback(ctx, b, update)

	opts := []echo.Option{
		echo.WithStartButtonText(loader.StartMsgDescText),
		echo.WithConfirmKeyboardText(loader.RegisterConfirmDescText),
		echo.WithCompleteText(loader.RegisterCompleteDescText),
		echo.WithConfirmFunction(h.proceedResult),
		echo.Questions([]echo.CollectItem{
			{
				FieldId:   "login",
				FieldName: "Логин: ",
				Content:   "Как мне к тебе обращаться?",
			},
			{
				FieldId:   "email",
				FieldName: "Почта: ",
				Content:   "Введи свой Email",
			},
		}),
	}
	c := echo.NewEcho(ctx, b, chatId, msgId, opts, h.dao.ActionRepo, "/start")
	h.echoComponent = *c
	c.StartCollect()
}

func (h *StartCommandHandler) ProceedMessage(ctx context.Context, b *bot.Bot, update *models.Update) {
	h.echoComponent.ProceedAnswer(ctx, b, update)
}

func (h *StartCommandHandler) proceedResult(result echo.Result) {
	h.dao.ActionRepo.SaveOrUpdate(result.ChatId, "StateConfirm", false, result.MsgId, "/start")
}

func (h *StartCommandHandler) GetName() string {
	return "/start"
}

func (h *StartCommandHandler) ClearStatus(update *models.Update) {
	chatId, _ := util.GetChatAndMsgId(update)
	tempMessageSlice[chatId] = TempUser{}
}

func (h *StartCommandHandler) buildKeyboard() models.ReplyMarkup {
	kb := &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "Да да...давай дальше", CallbackData: "start_callback"},
			},
		},
	}
	return kb
}

func (h *StartCommandHandler) getButtonsForStartKeyboard() [][]models.InlineKeyboardButton {
	return [][]models.InlineKeyboardButton{
		{
			{Text: "Да да...давай дальше"},
		},
	}
}

// TODO отловить ошибки + оптимизация
func (h *StartCommandHandler) callback(ctx context.Context, b *bot.Bot, update *models.Update) {
	message := util.GetChatMessage(update)
	user := tempMessageSlice[message.Chat.ID]
	var keyboard models.ReplyMarkup
	var text string
	var isCanEdit = true
	var isNeedSendMsg = true
	switch user.State {
	case StateDefault:
		//b.DeleteMessage(ctx, &bot.DeleteMessageParams{ChatID: message.Chat.ID, MessageID: message.ID})

		user.State = StateDrawHelloKeyboard
		text = loader.StartMsgDescText
		keyboard = h.buildKeyboard()
		isCanEdit = false
		h.dao.ActionRepo.SaveOrUpdate(message.Chat.ID, "StateDrawHelloKeyboard", false, message.ID, h.GetName())
	case StateDrawHelloKeyboard:
		user.State = StateAskEmail
		text = "Итак начнем...напиши мне свою почту"
		h.dao.ActionRepo.SaveOrUpdate(message.Chat.ID, "StateAskEmail", true, message.ID, h.GetName())
	case StateAskEmail:
		//msgToDelete = append(msgToDelete, message.ID)
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
			isNeedSendMsg = false
			isCanEdit = false
			fileData, errReadFile := loader.EnterEmailMistakeMem.ReadFile("files/ebat_ty_loh.jpg")
			if errReadFile != nil {
				fmt.Printf("error read file, %v\n", errReadFile)
			}

			b.SendPhoto(ctx, &bot.SendPhotoParams{
				ChatID:  message.Chat.ID,
				Photo:   &models.InputFileUpload{Filename: "ebat_ty_loh.jpg", Data: bytes.NewReader(fileData)},
				Caption: "Ну окэй поехали заново. Введи почту",
			})
			user.State = StateAskEmail
			h.dao.ActionRepo.SaveOrUpdate(message.Chat.ID, "StateAskEmail", true, message.ID, h.GetName())
		} else {
			//b.DeleteMessages(ctx, &bot.DeleteMessagesParams{ChatID: message.Chat.ID, MessageIDs: msgToDelete})
			accType := dto.TG
			// TODO отловаить ошибки -> RegisterUser
			h.authServerClient.RegisterUser(dto.CreateAccountRequest{
				AccountType: &accType,
				Email:       &user.Email,
				FirstName:   &user.FirstName,
				LastName:    &user.LastName,
				Login:       &user.Username,
				TelegramId:  &message.Chat.ID,
			})
			//TODO очистить диалог
			h.dao.ActionRepo.SaveOrUpdate(message.Chat.ID, "StateConfirm", false, message.ID, h.GetName())
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
	} else if isNeedSendMsg {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:      message.Chat.ID,
			Text:        text,
			ReplyMarkup: keyboard,
			ParseMode:   models.ParseModeHTML,
		})
		//msgToDelete = append(msgToDelete, sendMessage.ID)
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
