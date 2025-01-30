package command

import (
	"context"
	_ "embed"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/go-telegram/ui/dialog"
	"telegramCLient/internal/components"
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
	//go:embed tutorial_start_text.txt
	startText string

	//go:embed note_command_desc_text.txt
	noteCommandDescText string

	//go:embed profile_command_desc_text.txt
	profileCommandDescText string

	//go:embed menu_command_desc_text.txt
	menuCommandDescText string

	//go:embed expense_accounting_command_desc_text.txt
	expenseAccCommandDescText string

	dialogNodes = []dialog.Node{
		{ID: START_TEXT_ID, Text: startText, Keyboard: [][]dialog.Button{{{Text: CLOSE_TEXT, NodeID: SECOND_NODE_ID}, {Text: NEXT_TEXT, NodeID: THIRD_NODE_ID}}}},
		{ID: SECOND_NODE_ID, Text: END_TUTORIAL_MSG},
		{ID: THIRD_NODE_ID, Text: noteCommandDescText, Keyboard: [][]dialog.Button{{{Text: TO_BEGINNING_TEXT, NodeID: START_TEXT_ID}, {Text: NEXT_TEXT, NodeID: FOURTH_NODE_ID}}}},
		{ID: FOURTH_NODE_ID, Text: profileCommandDescText, Keyboard: [][]dialog.Button{{{Text: BACK_TEXT, NodeID: THIRD_NODE_ID}}, {{Text: NEXT_TEXT, NodeID: FIFTH_NODE_ID}}}},
		{ID: FIFTH_NODE_ID, Text: menuCommandDescText, Keyboard: [][]dialog.Button{{{Text: BACK_TEXT, NodeID: FOURTH_NODE_ID}}, {{Text: NEXT_TEXT, NodeID: SIX_NODE_ID}}}},
		{ID: SIX_NODE_ID, Text: expenseAccCommandDescText, Keyboard: [][]dialog.Button{{{Text: BACK_TEXT, NodeID: FIFTH_NODE_ID}}, {{Text: DONE_TEXT, NodeID: SECOND_NODE_ID}}}},
	}
)

type TutorialCommandHandler struct {
	Slider    components.Slider
	Paginator components.Paginator
	Dialog    components.DialogInline
}

func NewTutorialCommandHandler() *TutorialCommandHandler {
	return &TutorialCommandHandler{
		*components.NewSlider(),
		*components.NewPaginator(),
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
