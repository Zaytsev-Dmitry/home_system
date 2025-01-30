package components

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"strconv"
	"strings"
)

type Paginator struct {
	data              []string
	perPage           int
	currentPage       int
	pagesCount        int
	prefix            string
	callbackHandlerID string
	separator         string
	onError           OnErrorHandler
	closeButton       string
}

type OnErrorHandler func(err error)

const (
	cmdNop   = "nop"
	cmdStart = "start"
	cmdEnd   = "end"
	cmdClose = "close"
)

type Option func(p *Paginator)

func WithCloseButton(text string) Option {
	return func(p *Paginator) {
		p.closeButton = text
	}
}

func WithCallbackButton(text string, callbackData string) Option {
	return func(p *Paginator) {
		p.closeButton = text
	}
}

func NewPaginator(b *bot.Bot, data []string, perPage int, opts ...Option) *Paginator {
	p := &Paginator{
		data:        data,
		currentPage: 1,
		perPage:     perPage,
		prefix:      bot.RandomString(20),
		pagesCount:  len(data) / perPage,
		separator:   "\n\n\n",
	}
	for _, opt := range opts {
		opt(p)
	}

	//если количество доступных элементов/на кол-во элементов на странице не четное то инкрементим на 1 кол-во страниц
	if len(data)%p.perPage != 0 {
		p.pagesCount++
	}

	p.callbackHandlerID = b.RegisterHandler(bot.HandlerTypeCallbackQueryData, p.prefix, bot.MatchTypePrefix, p.callback)
	return p
}

func (p *Paginator) callback(ctx context.Context, b *bot.Bot, update *models.Update) {
	cmd := strings.TrimPrefix(update.CallbackQuery.Data, p.prefix)
	switch cmd {
	case cmdNop:
		p.callbackAnswer(ctx, b, update.CallbackQuery)
		return
	case cmdStart:
		if p.currentPage == 1 {
			p.callbackAnswer(ctx, b, update.CallbackQuery)
			return
		}
		p.currentPage = 1
	case cmdEnd:
		if p.currentPage == p.pagesCount {
			p.callbackAnswer(ctx, b, update.CallbackQuery)
			return
		}
		p.currentPage = p.pagesCount
	case cmdClose:
		b.UnregisterHandler(p.callbackHandlerID)

		_, errDelete := b.DeleteMessage(ctx, &bot.DeleteMessageParams{
			ChatID:    update.CallbackQuery.Message.Message.Chat.ID,
			MessageID: update.CallbackQuery.Message.Message.ID,
		})
		if errDelete != nil {
			p.onError(errDelete)
		}
		p.callbackAnswer(ctx, b, update.CallbackQuery)
		return
	default:
		page, _ := strconv.Atoi(cmd)
		p.currentPage = page
	}

	_, errEdit := b.EditMessageText(ctx, &bot.EditMessageTextParams{
		ChatID:          update.CallbackQuery.Message.Message.Chat.ID,
		MessageID:       update.CallbackQuery.Message.Message.ID,
		InlineMessageID: update.CallbackQuery.InlineMessageID,
		Text:            p.buildText(),
		ParseMode:       models.ParseModeMarkdown,
		ReplyMarkup:     p.buildKeyboard(),
	})
	if errEdit != nil {
		p.onError(errEdit)
	}

	p.callbackAnswer(ctx, b, update.CallbackQuery)
}

func (p *Paginator) callbackAnswer(ctx context.Context, b *bot.Bot, callbackQuery *models.CallbackQuery) {
	ok, err := b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: callbackQuery.ID,
	})
	if err != nil {
		p.onError(err)
		return
	}
	if !ok {
		p.onError(fmt.Errorf("callback answer failed"))
	}
}

func (p *Paginator) buildText() string {
	if len(p.data) <= p.perPage {
		return strings.Join(p.data, p.separator)
	}

	from := (p.currentPage - 1) * p.perPage
	to := from + p.perPage
	if to > len(p.data) {
		to = len(p.data)
	}

	return strings.Join(p.data[from:to], p.separator)
}

func (p *Paginator) buildKeyboard() models.InlineKeyboardMarkup {
	var row []models.InlineKeyboardButton
	if p.pagesCount <= 7 {
		if p.pagesCount > 1 {
			for i := 1; i <= p.pagesCount; i++ {
				callbackCommand := strconv.Itoa(i)
				buttonText := strconv.Itoa(i)
				if i == p.currentPage {
					buttonText = "( " + buttonText + " )"
				}

				row = append(row, models.InlineKeyboardButton{Text: buttonText, CallbackData: p.prefix + callbackCommand})
			}
		}
	} else {
		row = append(row, models.InlineKeyboardButton{Text: "\u00AB 1", CallbackData: p.prefix + cmdStart})

		startPage := p.calcStartPage()

		for i := startPage; i < startPage+5; i++ {
			callbackCommand := strconv.Itoa(i)
			buttonText := strconv.Itoa(i)
			if i > p.pagesCount {
				callbackCommand = cmdNop
				buttonText = " "
			}
			if i == p.currentPage {
				buttonText = "( " + buttonText + " )"
			}
			row = append(row, models.InlineKeyboardButton{Text: buttonText, CallbackData: p.prefix + callbackCommand})
		}
		row = append(row, models.InlineKeyboardButton{Text: strconv.Itoa(p.pagesCount) + " \u00BB", CallbackData: p.prefix + cmdEnd})
	}

	kb := models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{},
	}

	if len(row) > 0 {
		kb.InlineKeyboard = append(kb.InlineKeyboard, row)
	}

	if p.closeButton != "" {
		kb.InlineKeyboard = append(kb.InlineKeyboard, []models.InlineKeyboardButton{
			{Text: p.closeButton, CallbackData: p.prefix + cmdClose},
		})
	}

	return kb
}

// TODO какая то тупая логика
func (p *Paginator) calcStartPage() int {
	if p.pagesCount < 5 {
		return 1
	}
	if p.currentPage < 3 {
		return 1
	}
	if p.currentPage >= p.pagesCount-2 {
		return p.pagesCount - 4
	}
	return p.currentPage - 2
}

func (p *Paginator) CreateAndRun(ctx context.Context, b *bot.Bot, update *models.Update, addBtns []models.InlineKeyboardButton) (*models.Message, error) {
	keyboard := p.buildKeyboard()
	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, addBtns)

	params := &bot.SendMessageParams{
		ChatID:      update.CallbackQuery.Message.Message.Chat.ID,
		Text:        p.buildText(),
		ParseMode:   models.ParseModeMarkdown,
		ReplyMarkup: keyboard,
	}
	return b.SendMessage(ctx, params)
}
