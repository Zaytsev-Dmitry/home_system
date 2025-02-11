package command

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"telegramCLient/external"
	"telegramCLient/internal/components/paginator"
	"telegramCLient/internal/handler/loader"
	"telegramCLient/util"
)

type NoteCommandHandler struct {
	noteBackClient *external.NoteBackendClient
}

func NewNoteCommandHandler(client *external.NoteBackendClient) *NoteCommandHandler {
	return &NoteCommandHandler{
		noteBackClient: client,
	}
}

func (h *NoteCommandHandler) Init() []bot.Option {
	return []bot.Option{
		bot.WithMessageTextHandler("/notes", bot.MatchTypeExact, h.showNotesKeyboardCallback),
		bot.WithCallbackQueryDataHandler("open_notes", bot.MatchTypeExact, h.showNotesKeyboardCallback),
		bot.WithCallbackQueryDataHandler("add_note", bot.MatchTypeExact, h.addNoteCallback),
		bot.WithCallbackQueryDataHandler("delete_note", bot.MatchTypeExact, h.deleteNoteCallback),
		bot.WithCallbackQueryDataHandler("show_all_notes", bot.MatchTypeExact, h.showAllNoteCallback),
		bot.WithCallbackQueryDataHandler("show_note_by_name", bot.MatchTypeExact, h.showNoteByNameCallback),
	}
}

func (h *NoteCommandHandler) StartCommand(ctx context.Context, b *bot.Bot, update *models.Update) {

}

func (h *NoteCommandHandler) ProceedMessage(ctx context.Context, b *bot.Bot, update *models.Update) {

}

func (h *NoteCommandHandler) GetName() string {
	return "/notes"
}

func (h *NoteCommandHandler) ClearStatus(update *models.Update) {
}

func (h *NoteCommandHandler) AddToDelete(msg int) {

}

// TODO отловаить ошибки
func (h *NoteCommandHandler) addNoteCallback(ctx context.Context, b *bot.Bot, update *models.Update) {

}

// TODO отловаить ошибки
func (h *NoteCommandHandler) deleteNoteCallback(ctx context.Context, b *bot.Bot, update *models.Update) {

}

func (h *NoteCommandHandler) showAllNoteCallback(ctx context.Context, b *bot.Bot, update *models.Update) {
	data := h.noteBackClient.GetAllNotesByAccount(update.CallbackQuery.From.ID)
	var stringData []string
	for _, value := range *data.Objects {
		text := fmt.Sprintf(
			loader.NoteCommandContentText,
			*value.Id,
			*value.Name,
			*value.Description,
		)
		stringData = append(stringData, text)
	}
	paginator.NewPaginator().CreateAndRun(ctx, b, update, stringData, 5, "Закрыть ❌")
}

func (h *NoteCommandHandler) showNoteByNameCallback(ctx context.Context, b *bot.Bot, update *models.Update) {

}

// TODO добавить edit msg если приходят из меню
func (h *NoteCommandHandler) showNotesKeyboardCallback(ctx context.Context, b *bot.Bot, update *models.Update) {
	chatId, _ := util.GetChatAndMsgId(update)
	b.SendMessage(
		ctx,
		&bot.SendMessageParams{
			ChatID:      chatId,
			Text:        "Выбери действие",
			ReplyMarkup: h.buildNotesKeyboard(),
		})
}

func (h *NoteCommandHandler) buildNotesKeyboard() models.ReplyMarkup {
	kb := &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "Добавить запись 📄", CallbackData: "add_note"},
			},
			{
				{Text: "Удалить записку ❌", CallbackData: "delete_note"},
			},
			{
				{Text: "Просмотреть все записки 🗄️", CallbackData: "show_all_notes"},
			},
			{
				{Text: "Поиск записки по названию 🔎", CallbackData: "show_note_by_name"},
			},
			{
				{Text: "Назад к меню 🤙", CallbackData: "open_menu"},
			},
		},
	}
	return kb
}
