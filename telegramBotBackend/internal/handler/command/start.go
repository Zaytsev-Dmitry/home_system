package command

import (
	"context"
	authSpec "github.com/Zaytsev-Dmitry/home_system_open_api/authServerBackend"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"telegramCLient/external"
	"telegramCLient/internal/components/echo"
	"telegramCLient/internal/dao"
	"telegramCLient/internal/handler/loader"
	"telegramCLient/util"
)

type StartCommandHandler struct {
	dao              dao.TelegramBotDao
	authServerClient external.AuthServerClient
	echoComponent    *echo.Echo
}

func NewStartCommandHandler(d dao.TelegramBotDao, serverClient external.AuthServerClient) *StartCommandHandler {
	return &StartCommandHandler{dao: d, authServerClient: serverClient}
}

func (h *StartCommandHandler) Init() []bot.Option {
	return []bot.Option{
		bot.WithMessageTextHandler("/start", bot.MatchTypeExact, h.StartCommand),
	}
}

func (h *StartCommandHandler) StartCommand(ctx context.Context, b *bot.Bot, update *models.Update) {
	chatId, msgId := util.GetChatAndMsgId(update)
	////TODO проверить зареган ли пользак или нет
	opts := []echo.Option{
		echo.WithControlMessage(h.dao.ActionRepo),
		echo.WithStartButtonText(loader.StartCommandText),
		echo.WithConfirmKeyboardText(loader.RegisterConfirmText),
		echo.WithCompleteText(loader.RegisterCompleteText),
		echo.WithConfirmFunction(h.proceedResult),
		echo.Questions([]echo.CollectItem{
			{
				FieldId:   "username",
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
	c := echo.NewEcho(ctx, b, chatId, msgId, opts, h.dao.ActionRepo, h.GetName())
	h.echoComponent = c
	c.StartCollect()
}

func (h *StartCommandHandler) ProceedMessage(ctx context.Context, b *bot.Bot, update *models.Update) {
	h.echoComponent.ProceedAnswer(ctx, b, update)
}

func (h *StartCommandHandler) proceedResult(result echo.Result) {
	h.dao.ActionRepo.SaveOrUpdate(result.ChatId, "Done", false, result.MsgId, h.GetName())
	accType := authSpec.TG
	request := authSpec.CreateAccountRequest{
		AccountType:      &accType,
		TelegramId:       &result.ChatId,
		FirstName:        &result.UserFirstName,
		LastName:         &result.UserLastname,
		TelegramUserName: &result.UserTGName,
	}
	for i, answer := range result.Answers {
		if answer.FieldId == "username" {
			request.Username = &result.Answers[i].Content
		}
		if answer.FieldId == "email" {
			request.Email = &result.Answers[i].Content
		}
	}
	//TODO если не смог зарегать то надо отправлять ошибку и как бы банить выполнение команды
	h.authServerClient.RegisterUser(request)
}

func (h *StartCommandHandler) GetName() string {
	return "/start"
}

func (h *StartCommandHandler) ClearStatus(update *models.Update) {
	h.echoComponent.Clear(update)
}

func (h *StartCommandHandler) AddToDelete(msg int) {
	h.echoComponent.AddToDelete(msg)
}
