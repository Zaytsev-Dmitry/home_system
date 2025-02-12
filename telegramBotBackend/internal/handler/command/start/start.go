package start

import (
	"context"
	authSpec "github.com/Zaytsev-Dmitry/home_system_open_api/authServerBackend"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"telegramCLient/external"
	"telegramCLient/internal/components/echo"
	"telegramCLient/internal/dao"
	"telegramCLient/internal/handler/command"
	"telegramCLient/internal/handler/loader"
	"telegramCLient/internal/storage"
)

type StartCommand struct {
	dao               dao.TelegramBotDao
	authServerClient  *external.AuthServerClient
	component         *echo.Echo
	messageStorage    storage.Storage
	ctx               context.Context
	bot               *bot.Bot
	callbackHandlerID string
	action            command.Action
}

func NewStartCommand(action command.Action, st storage.Storage, bot *bot.Bot, ctx context.Context, d dao.TelegramBotDao, serverClient *external.AuthServerClient) *StartCommand {
	s := &StartCommand{
		dao:              d,
		authServerClient: serverClient,
		messageStorage:   st,
		bot:              bot,
		ctx:              ctx,
		action:           action,
	}
	textMeta := echo.TextMeta{
		ConfirmText: loader.RegisterConfirmText,
		StartText:   loader.StartCommandText,
	}
	options := []echo.Option{
		echo.WithMessageStorage(st),
	}
	s.component = echo.NewEcho(bot, s.getQuestions(), s.proceedResult, s.setUserInput, textMeta, options)
	return s
}

func (s *StartCommand) RegisterHandler() {
	s.callbackHandlerID = s.bot.RegisterHandler(bot.HandlerTypeMessageText, s.GetName(), bot.MatchTypeExact, s.callback)
}

func (s *StartCommand) proceedUserAnswer(ctx context.Context, b *bot.Bot, update *models.Update) {
	s.component.ProceedUserAnswer(ctx, b, update)
}

func (s *StartCommand) setUserInput(userInput bool, chatId int64) {
	s.action.Log(chatId, s.GetName(), userInput, true)
}

func (s *StartCommand) proceedResult(result echo.Result) {
	accType := authSpec.TG
	request := authSpec.CreateAccountRequest{
		AccountType:      &accType,
		TelegramId:       &result.ChatId,
		FirstName:        &result.UserFirstName,
		LastName:         &result.UserLastname,
		TelegramUserName: &result.Username,
	}
	for i, answer := range result.Question {
		if answer.FieldId == "username" {
			request.Username = &result.Question[i].Answer
		}
		if answer.FieldId == "email" {
			request.Email = &result.Question[i].Answer
		}
	}
	//TODO если не смог зарегать то надо отправлять ошибку и как бы банить выполнение команды
	_, err := s.authServerClient.RegisterUser(request)
	if err != nil {
		//TODO отправить ошибку в чат
	} else {
		s.bot.SendMessage(
			s.ctx, &bot.SendMessageParams{
				Text:      loader.RegisterCompleteText,
				ChatID:    result.ChatId,
				ParseMode: models.ParseModeHTML,
			})
		s.bot.DeleteMessages(
			s.ctx, &bot.DeleteMessagesParams{
				ChatID:     result.ChatId,
				MessageIDs: result.MessagesIds[0:len(result.MessagesIds)],
			},
		)
		s.messageStorage.ClearAll(result.ChatId)
	}
}

func (s *StartCommand) ProceedUserAnswer(ctx context.Context, b *bot.Bot, update *models.Update) {
	s.component.ProceedUserAnswer(ctx, b, update)
}

func (s *StartCommand) GetName() string {
	return "/start"
}

func (s *StartCommand) getQuestions() []echo.CollectItem {
	return []echo.CollectItem{
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
	}
}
func (s *StartCommand) ClearState(chatId int64) {
	s.component.ClearState(chatId)
	allMsg := s.messageStorage.GetAll(chatId)
	var idsToDel []int
	for _, msg := range allMsg {
		idsToDel = append(idsToDel, msg.Id)
	}
	s.bot.DeleteMessages(s.ctx, &bot.DeleteMessagesParams{
		ChatID:     chatId,
		MessageIDs: idsToDel,
	})
}
