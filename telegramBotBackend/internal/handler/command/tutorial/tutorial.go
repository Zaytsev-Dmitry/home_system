package tutorial

import (
	"context"
	_ "embed"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/go-telegram/ui/dialog"
	dialog2 "telegramCLient/internal/components/dialog"
	"telegramCLient/internal/handler/loader"
	"telegramCLient/internal/storage"
)

const (
	END_TUTORIAL_MSG  = "*Теперь ты знаешь что я умею и ты готов к работе*"
	CLOSE_TEXT        = "Закрыть"
	NEXT_TEXT         = "Дальше"
	BACK_TEXT         = "Назад"
	DONE_TEXT         = "Я все понял"
	START_TEXT_ID     = "start"
	TO_BEGINNING_TEXT = "В начало"
	SECOND_NODE_ID    = "2"
	THIRD_NODE_ID     = "3"
	FOURTH_NODE_ID    = "4"
	FIFTH_NODE_ID     = "5"
	SIX_NODE_ID       = "6"
)

var (
	dialogNodes = []dialog.Node{
		{ID: START_TEXT_ID, Text: loader.TutorialStartText, Keyboard: [][]dialog.Button{{{Text: CLOSE_TEXT, NodeID: SECOND_NODE_ID}, {Text: NEXT_TEXT, NodeID: THIRD_NODE_ID}}}},
		{ID: SECOND_NODE_ID, Text: END_TUTORIAL_MSG},
		{ID: THIRD_NODE_ID, Text: loader.NoteCommandText, Keyboard: [][]dialog.Button{{{Text: TO_BEGINNING_TEXT, NodeID: START_TEXT_ID}, {Text: NEXT_TEXT, NodeID: FOURTH_NODE_ID}}}},
		{ID: FOURTH_NODE_ID, Text: loader.ProfileCommandText, Keyboard: [][]dialog.Button{{{Text: BACK_TEXT, NodeID: THIRD_NODE_ID}}, {{Text: NEXT_TEXT, NodeID: FIFTH_NODE_ID}}}},
		{ID: FIFTH_NODE_ID, Text: loader.MenuCommandText, Keyboard: [][]dialog.Button{{{Text: BACK_TEXT, NodeID: FOURTH_NODE_ID}}, {{Text: NEXT_TEXT, NodeID: SIX_NODE_ID}}}},
		{ID: SIX_NODE_ID, Text: loader.ExpenseAccountingCommandText, Keyboard: [][]dialog.Button{{{Text: BACK_TEXT, NodeID: FIFTH_NODE_ID}}, {{Text: DONE_TEXT, NodeID: SECOND_NODE_ID}}}},
	}
)

type TutorialCommand struct {
	component         dialog2.DialogInline
	messageStorage    storage.Storage
	ctx               context.Context
	bot               *bot.Bot
	callbackHandlerID string
}

func NewTutorialCommand(st storage.Storage, bot *bot.Bot, ctx context.Context) *TutorialCommand {
	return &TutorialCommand{
		component:      *dialog2.NewDialogInline(),
		messageStorage: st,
		bot:            bot,
		ctx:            ctx,
	}
}

func (t *TutorialCommand) RegisterHandler() {
	t.callbackHandlerID = t.bot.RegisterHandler(bot.HandlerTypeMessageText, t.GetName(), bot.MatchTypeExact, t.callback)
}

func (t *TutorialCommand) ProceedUserAnswer(ctx context.Context, b *bot.Bot, update *models.Update) {
}

func (t *TutorialCommand) GetName() string {
	return "/tutorial"
}
