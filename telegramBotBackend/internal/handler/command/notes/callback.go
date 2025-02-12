package notes

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"telegramCLient/internal/components/paginator"
	"telegramCLient/internal/handler/loader"
	"telegramCLient/internal/storage"
	"telegramCLient/util"
)

func (n *NoteCommand) addNoteCallback(ctx context.Context, b *bot.Bot, update *models.Update) {
	//TODO вызывать logcommon action
	n.component.Collect(ctx, b, update)
}

// TODO если ошибки то обработать
func (n *NoteCommand) showAllNotesCallback(ctx context.Context, b *bot.Bot, update *models.Update) {
	//TODO вызывать logcommon action
	data := n.noteBackClient.GetAllNotesByAccount(update.CallbackQuery.From.ID)
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

func (n *NoteCommand) showNotesKeyboardCallback(ctx context.Context, b *bot.Bot, update *models.Update) {
	chatId, _ := util.GetChatAndMsgId(update)
	message, _ := b.SendMessage(
		ctx,
		&bot.SendMessageParams{
			ChatID:      chatId,
			Text:        "Выбери действие",
			ReplyMarkup: n.buildNotesKeyboard(),
		})
	n.messageStorage.Add(chatId, *storage.NewMessage(message.ID, message.Text, 0, storage.BOT))
	n.LogCommandAction(chatId, "start")
}
