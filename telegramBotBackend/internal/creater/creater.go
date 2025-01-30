package creater

import (
	"fmt"
	"github.com/go-telegram/bot"
	"log"
	"telegramCLient/config"
	"telegramCLient/internal/handler/command"
)

const (
	START_HANDLER    string = "StartCommandHandler"
	REGISTER_HANDLER string = "RegisterCommandHandler"
	TUTORIAL_HANDLER string = "TutorialCommandHandler"
	MENU_HANDLER     string = "MenuCommandHandler"
	NOTE_HANDLER     string = "NoteCommandHandler"
	PROFILE_HANDLER  string = "ProfileCommandHandler"
)

type HandlerCreater struct {
	Config                *config.AppConfig
	tempMessageCollection map[int64][]int
}

func CreateHandlerStarter(conf *config.AppConfig, tempMessage map[int64][]int) *HandlerCreater {
	return &HandlerCreater{
		Config:                conf,
		tempMessageCollection: tempMessage,
	}
}

// TODO отловаить ошибки
func (h *HandlerCreater) CreateHandlers() []bot.Option {
	var result = []bot.Option{}
	for i, value := range h.Config.HandlersToInit {
		log.Println(fmt.Sprintf("CreateHandlers handler : %s. With order: %x", value, i+1))
		switch value {
		case START_HANDLER:
			{
				result = append(result, command.NewStartCommandHandler(h.tempMessageCollection).Init()...)
			}
		case REGISTER_HANDLER:
			{
				result = append(result, command.NewRegisterUserCommandHandler(h.Config.AuthServerUrl, h.tempMessageCollection).Init()...)
			}

		case TUTORIAL_HANDLER:
			{
				result = append(result, command.NewTutorialCommandHandler().Init()...)
			}
		case MENU_HANDLER:
			{
				result = append(result, command.NewMenuCommandHandler().Init()...)
			}
		case NOTE_HANDLER:
			{
				result = append(result, command.NewNoteCommandHandler().Init()...)
			}

		case PROFILE_HANDLER:
			{
				result = append(result, command.NewProfileCommandHandler().Init()...)
			}
		}
	}
	return result
}
