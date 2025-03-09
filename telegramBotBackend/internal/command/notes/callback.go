package notes

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"telegramCLient/internal/command/loader"
	"telegramCLient/internal/components/paginator"
	"telegramCLient/internal/storage"
	"telegramCLient/util"
)

func (n *NoteCommand) addNoteCallback(ctx context.Context, b *bot.Bot, update *models.Update) {
	message := util.GetChatMessage(update)
	n.action.Log(message.Chat.ID, n.GetName(), false, true)
	n.component.Collect(ctx, b, update)
}

// TODO если ошибки то обработать
func (n *NoteCommand) showAllNotesCallback(ctx context.Context, b *bot.Bot, update *models.Update) {
	message := util.GetChatMessage(update)
	n.action.Log(message.Chat.ID, n.GetName(), false, true)

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
	n.action.Log(message.Chat.ID, n.GetName(), false, true)
	n.messageStorage.Add(chatId, *storage.NewMessage(message.ID, message.Text, storage.BOT))
}
