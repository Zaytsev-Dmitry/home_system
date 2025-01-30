package command

import (
	"context"
	_ "embed"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/go-telegram/ui/dialog"
	"telegramCLient/internal/components"
	"telegramCLient/internal/handler/loader"
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
		{ID: START_TEXT_ID, Text: loader.StartText, Keyboard: [][]dialog.Button{{{Text: CLOSE_TEXT, NodeID: SECOND_NODE_ID}, {Text: NEXT_TEXT, NodeID: THIRD_NODE_ID}}}},
		{ID: SECOND_NODE_ID, Text: END_TUTORIAL_MSG},
		{ID: THIRD_NODE_ID, Text: loader.NoteCommandDescText, Keyboard: [][]dialog.Button{{{Text: TO_BEGINNING_TEXT, NodeID: START_TEXT_ID}, {Text: NEXT_TEXT, NodeID: FOURTH_NODE_ID}}}},
		{ID: FOURTH_NODE_ID, Text: loader.ProfileCommandDescText, Keyboard: [][]dialog.Button{{{Text: BACK_TEXT, NodeID: THIRD_NODE_ID}}, {{Text: NEXT_TEXT, NodeID: FIFTH_NODE_ID}}}},
		{ID: FIFTH_NODE_ID, Text: loader.MenuCommandDescText, Keyboard: [][]dialog.Button{{{Text: BACK_TEXT, NodeID: FOURTH_NODE_ID}}, {{Text: NEXT_TEXT, NodeID: SIX_NODE_ID}}}},
		{ID: SIX_NODE_ID, Text: loader.ExpenseAccCommandDescText, Keyboard: [][]dialog.Button{{{Text: BACK_TEXT, NodeID: FIFTH_NODE_ID}}, {{Text: DONE_TEXT, NodeID: SECOND_NODE_ID}}}},
	}
)

type TutorialCommandHandler struct {
	Dialog components.DialogInline
}

func NewTutorialCommandHandler() *TutorialCommandHandler {
	return &TutorialCommandHandler{
		*components.NewDialogInline(),
	}
}

func (c *TutorialCommandHandler) Init() []bot.Option {
	return []bot.Option{
		bot.WithMessageTextHandler("/tutorial", bot.MatchTypeExact, c.tutorialCallback),
	}
}

func (c *TutorialCommandHandler) tutorialCallback(ctx context.Context, b *bot.Bot, update *models.Update) {
	c.Dialog.CreateAndRun(dialogNodes, ctx, b, update)
}
