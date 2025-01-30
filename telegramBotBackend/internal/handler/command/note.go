package command

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"telegramCLient/internal/components"
)

var (
	fakeData = []string{
		"*1* Lorem ipsum dolor sit amet, consectetur adipiscing elit",
		"*2* Cras faucibus ipsum vel sodales tincidunt",
		"*3* Nulla bibendum lacus vitae arcu rutrum, quis accumsan dolor auctor",
		"*4* Morbi non mi nec nulla rutrum aliquet",
		"*5* Curabitur quis leo facilisis, vulputate sem id, euismod mauris",
		"*6* Sed condimentum tellus in diam dapibus, in euismod nisi ornare",
		"*7* Quisque ut neque congue, blandit orci vitae, viverra mi",
		"*8* Suspendisse porttitor erat in hendrerit pulvinar",
		"*9* Suspendisse cursus arcu placerat auctor vehicula",
		"*10* Phasellus tempus nisi a urna luctus aliquam",
	}
)

type NoteCommandHandler struct {
}

func NewNoteCommandHandler() *NoteCommandHandler {
	return &NoteCommandHandler{}
}

func (h *NoteCommandHandler) Init() []bot.Option {
	return []bot.Option{
		bot.WithMessageTextHandler("/notes", bot.MatchTypeExact, h.showNotesKeyboardCallback),
		bot.WithCallbackQueryDataHandler("add_note", bot.MatchTypeExact, h.addNoteCallback),
		bot.WithCallbackQueryDataHandler("delete_note", bot.MatchTypeExact, h.deleteNoteCallback),
		bot.WithCallbackQueryDataHandler("show_all_notes", bot.MatchTypeExact, h.showAllNoteCallback),
		bot.WithCallbackQueryDataHandler("show_note_by_name", bot.MatchTypeExact, h.showNoteByNameCallback),
	}
}

// TODO отловаить ошибки
func (h *NoteCommandHandler) addNoteCallback(ctx context.Context, b *bot.Bot, update *models.Update) {

}

// TODO отловаить ошибки
func (h *NoteCommandHandler) deleteNoteCallback(ctx context.Context, b *bot.Bot, update *models.Update) {

}

func (h *NoteCommandHandler) showAllNoteCallback(ctx context.Context, b *bot.Bot, update *models.Update) {
	components.NewPaginator().CreateAndRun(ctx, b, update, fakeData, 5)
}

func (h *NoteCommandHandler) showNoteByNameCallback(ctx context.Context, b *bot.Bot, update *models.Update) {

}

// TODO отловаить ошибки
func (h *NoteCommandHandler) showNotesKeyboardCallback(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        fmt.Sprintf("Выбери действие"),
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
		},
	}
	return kb
}
