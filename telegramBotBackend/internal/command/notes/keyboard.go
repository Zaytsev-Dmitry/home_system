package notes

import "github.com/go-telegram/bot/models"

func (h *NoteCommand) buildNotesKeyboard() models.ReplyMarkup {
	return &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "Добавить запись 📄", CallbackData: "add_note"},
			},
			{
				//{Text: "Удалить записку ❌", CallbackData: "delete_note"},
			},
			{
				{Text: "Просмотреть все записки 🗄️", CallbackData: "show_all_notes"},
			},
			{
				//{Text: "Поиск записки по названию 🔎", CallbackData: "show_note_by_name"},
			},
			{
				{Text: "Назад к меню 🤙", CallbackData: "open_menu"},
			},
		},
	}
}
