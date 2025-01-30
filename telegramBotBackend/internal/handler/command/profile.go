package command

import "github.com/go-telegram/bot"

type ProfileCommandHandler struct {
}

func NewProfileCommandHandler() *ProfileCommandHandler {
	return &ProfileCommandHandler{}
}

func (h *ProfileCommandHandler) Init() []bot.Option {
	return []bot.Option{}
}
