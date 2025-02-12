package notes

import (
	"context"
	noteSpec "github.com/Zaytsev-Dmitry/home_system_open_api/noteServerBackend"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"telegramCLient/external"
	"telegramCLient/internal/components/echo"
	"telegramCLient/internal/dao"
	"telegramCLient/internal/handler/command"
	"telegramCLient/internal/handler/loader"
	"telegramCLient/internal/storage"
)

type NoteCommand struct {
	noteBackClient    *external.NoteBackendClient
	component         *echo.Echo
	dao               dao.TelegramBotDao
	messageStorage    storage.Storage
	ctx               context.Context
	bot               *bot.Bot
	callbackHandlerID []string
	action            command.Action
}

func NewNotesCommand(action command.Action, st storage.Storage, bot *bot.Bot, ctx context.Context, d dao.TelegramBotDao, client *external.NoteBackendClient) *NoteCommand {
	n := &NoteCommand{
		noteBackClient: client,
		dao:            d,
		messageStorage: st,
		ctx:            ctx,
		bot:            bot,
		action:         action,
	}
	textMeta := echo.TextMeta{
		ConfirmText: loader.AddNoteConfirmCommandText,
		StartText:   loader.AddNoteStartCommandText,
	}
	options := []echo.Option{
		echo.WithMessageStorage(st),
	}
	n.component = echo.NewEcho(bot, n.getQuestions(), n.proceedResult, n.setUserInput, textMeta, options)
	return n
}

func (n *NoteCommand) RegisterHandler() {
	n.callbackHandlerID = append(n.callbackHandlerID, n.bot.RegisterHandler(bot.HandlerTypeMessageText, "/notes", bot.MatchTypeExact, n.showNotesKeyboardCallback))
	n.callbackHandlerID = append(n.callbackHandlerID, n.bot.RegisterHandler(bot.HandlerTypeCallbackQueryData, "add_note", bot.MatchTypeExact, n.addNoteCallback))
	n.callbackHandlerID = append(n.callbackHandlerID, n.bot.RegisterHandler(bot.HandlerTypeCallbackQueryData, "show_all_notes", bot.MatchTypeExact, n.showAllNotesCallback))
}

func (n *NoteCommand) ProceedUserAnswer(ctx context.Context, b *bot.Bot, update *models.Update) {
	n.component.ProceedUserAnswer(ctx, b, update)
}

func (n *NoteCommand) setUserInput(userInput bool, chatId int64) {
	n.action.Log(chatId, n.GetName(), userInput, true)
}

func (s *NoteCommand) proceedResult(result echo.Result) {
	request := noteSpec.CreateNoteRequest{
		TgId: &result.ChatId,
	}
	for i, answer := range result.Question {
		if answer.FieldId == "name" {
			request.Name = &result.Question[i].Answer
		}
		if answer.FieldId == "content" {
			request.Description = &result.Question[i].Answer
		}
		if answer.FieldId == "link" {
			request.Link = &result.Question[i].Answer
		}
	}
	s.noteBackClient.Save(request)

	s.bot.DeleteMessages(
		s.ctx, &bot.DeleteMessagesParams{
			ChatID:     result.ChatId,
			MessageIDs: result.MessagesIds[0:len(result.MessagesIds)],
		},
	)
	s.messageStorage.ClearAll(result.ChatId)

	s.bot.SendMessage(
		s.ctx,
		&bot.SendMessageParams{
			ChatID:      result.ChatId,
			Text:        "Выбери действие",
			ReplyMarkup: s.buildNotesKeyboard(),
		})
}

func (n *NoteCommand) GetName() string {
	return "/notes"
}

func (n *NoteCommand) getQuestions() []echo.CollectItem {
	return []echo.CollectItem{
		{
			FieldId:   "name",
			FieldName: "Название: ",
			Content:   "какое название записки?",
		},
		{
			FieldId:   "content",
			FieldName: "Содержимое: ",
			Content:   "напиши содержимое",
		},
		{
			FieldId:   "link",
			FieldName: "Ссылка: ",
			Content:   "какая ссылка",
		},
	}
}
func (n *NoteCommand) ClearState(chatId int64) {
	n.component.ClearState(chatId)
	allMsg := n.messageStorage.GetAll(chatId)
	var idsToDel []int
	for _, msg := range allMsg {
		idsToDel = append(idsToDel, msg.Id)
	}
	n.bot.DeleteMessages(n.ctx, &bot.DeleteMessagesParams{
		ChatID:     chatId,
		MessageIDs: idsToDel,
	})
}
