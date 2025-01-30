package command

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/go-telegram/ui/dialog"
	"telegramCLient/internal/components"
)

const FIRST_PAGE_CONTENT = "*На данный момент есть несколько доступных тебе команд*\n\n" +
	"\u2714 *\"notes\"* \u2011 записки\n\n" +
	"\u2714 *\"profile\"* \u2011 информация о твоем аккаунте\n\n" +
	"\u2714 *\"menu\"* \u2011 вызов меню\n\n" +
	"\u2714 *\"expense accounting\"* \u2011 учет совместных расходов\n\n" +
	"нажми *\"Дальше\"* чтобы узнать более подробно о каждой команде и функционале"

const NOTES_COMMAND_INFO_PAGE_CONTENT = "" +
	" Команда *\"notes\"* или проще говоря записки \n\n" +
	"Ты можешь *сохранять удалять редактированить просмотривать записки* \n\n" +
	"Если вдруг надо что то сохранить смело пользуйся"

const PROFILE_COMMAND_INFO_PAGE_CONTENT = "" +
	"Команда *\"profile\"* это все что касается твоего профиля \n\n" +
	"*Там хранится информация о твоем пользователе *\n\n" +
	"Увы но сейчас доступен только просмотр без редактирования"

const MENU_COMMAND_INFO_PAGE_CONTENT = "" +
	"Команда *\"menu\"* Это вызов контекстного меню \n\n"

const EXPENSE_ACCOUNTING__INFO_PAGE_CONTENT = "" +
	"Команда *\"expense accounting\"* Рано радуешься я еще не сделал это говно \n\n"

const END_TUTORIAL_MSG = "*Теперь ты знаешь что я умею и ты готов к работе*"
const CLOSE_TEXT = "Закрыть"
const NEXT_TEXT = "Дальше"
const BACK_TEXT = "Назад"
const DONE_TEXT = "Я все понял"
const START_TEXT_ID = "start"
const TO_BEGINNING_TEXT = "В начало"
const SECOND_NODE_ID = "2"
const THIRD_NODE_ID = "3"
const FOURTH_NODE_ID = "4"
const FIFTH_NODE_ID = "5"
const SIX_NODE_ID = "6"

var (
	dialogNodes = []dialog.Node{
		{ID: START_TEXT_ID, Text: FIRST_PAGE_CONTENT, Keyboard: [][]dialog.Button{{{Text: CLOSE_TEXT, NodeID: SECOND_NODE_ID}, {Text: NEXT_TEXT, NodeID: THIRD_NODE_ID}}}},
		{ID: SECOND_NODE_ID, Text: END_TUTORIAL_MSG},
		{ID: THIRD_NODE_ID, Text: NOTES_COMMAND_INFO_PAGE_CONTENT, Keyboard: [][]dialog.Button{{{Text: TO_BEGINNING_TEXT, NodeID: START_TEXT_ID}, {Text: NEXT_TEXT, NodeID: FOURTH_NODE_ID}}}},
		{ID: FOURTH_NODE_ID, Text: PROFILE_COMMAND_INFO_PAGE_CONTENT, Keyboard: [][]dialog.Button{{{Text: BACK_TEXT, NodeID: THIRD_NODE_ID}}, {{Text: NEXT_TEXT, NodeID: FIFTH_NODE_ID}}}},
		{ID: FIFTH_NODE_ID, Text: MENU_COMMAND_INFO_PAGE_CONTENT, Keyboard: [][]dialog.Button{{{Text: BACK_TEXT, NodeID: FOURTH_NODE_ID}}, {{Text: NEXT_TEXT, NodeID: SIX_NODE_ID}}}},
		{ID: SIX_NODE_ID, Text: EXPENSE_ACCOUNTING__INFO_PAGE_CONTENT, Keyboard: [][]dialog.Button{{{Text: BACK_TEXT, NodeID: FIFTH_NODE_ID}}, {{Text: DONE_TEXT, NodeID: SECOND_NODE_ID}}}},
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
