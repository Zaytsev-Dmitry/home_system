package echo

import (
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"telegramCLient/internal/dao/repository/intefraces"
)

type Option func(e *Echo)

func WithStartButtonText(text string) Option {
	return func(p *Echo) {
		prefix := p.prefix + startKeyboardCallback
		keyboard := &models.InlineKeyboardMarkup{
			InlineKeyboard: [][]models.InlineKeyboardButton{
				{
					{Text: "Да да...давай дальше ⏭️", CallbackData: prefix},
				},
			},
		}
		p.startKeyboardText = text
		p.startKeyboard = keyboard
		handlerUid := p.bot.RegisterHandler(bot.HandlerTypeCallbackQueryData, prefix, bot.MatchTypePrefix, p.callback)
		p.callbackHandlerIDs = append(p.callbackHandlerIDs, handlerUid)
	}
}

func WithCompleteText(text string) Option {
	return func(p *Echo) {
		p.completeText = text
	}
}

func WithControlMessage(repo intefraces.ActionRepository) Option {
	return func(p *Echo) {
		p.actionRepo = repo
	}
}

func WithConfirmFunction(confirmFunc func(result Result)) Option {
	return func(p *Echo) {
		p.confirmCallbackFunction = confirmFunc
	}
}

func WithConfirmKeyboardText(text string) Option {
	return func(p *Echo) {
		yesCallback := p.prefix + confirmCallbackYes
		noCallback := p.prefix + confirmCallbackNo
		keyboard := &models.InlineKeyboardMarkup{
			InlineKeyboard: [][]models.InlineKeyboardButton{
				{
					{Text: "Да ✅", CallbackData: yesCallback},
					{Text: "Нет ❌", CallbackData: noCallback},
				},
			},
		}
		p.confirmKeyboard = keyboard
		p.confirmKeyboardText = text
		yesUid := p.bot.RegisterHandler(bot.HandlerTypeCallbackQueryData, yesCallback, bot.MatchTypePrefix, p.callback)
		noUid := p.bot.RegisterHandler(bot.HandlerTypeCallbackQueryData, noCallback, bot.MatchTypePrefix, p.callback)

		p.callbackHandlerIDs = append(p.callbackHandlerIDs, yesUid)
		p.callbackHandlerIDs = append(p.callbackHandlerIDs, noUid)
	}
}

func Questions(questions []CollectItem) Option {
	return func(p *Echo) {
		p.questions = questions
	}
}
