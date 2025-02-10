package creater

import (
	"fmt"
	"github.com/go-telegram/bot"
	"log"
	"telegramCLient/config"
	"telegramCLient/external"
	"telegramCLient/internal/dao"
	"telegramCLient/internal/handler/command"
)

const (
	START_HANDLER    string = "StartCommandHandler"
	TUTORIAL_HANDLER string = "TutorialCommandHandler"
	MENU_HANDLER     string = "MenuCommandHandler"
	NOTE_HANDLER     string = "NoteCommandHandler"
	PROFILE_HANDLER  string = "ProfileCommandHandler"
)

type HandlerCreater struct {
	Config *config.AppConfig
}

func CreateHandlerStarter(conf *config.AppConfig) *HandlerCreater {
	return &HandlerCreater{
		Config: conf,
	}
}

// TODO отловаить ошибки
func (h *HandlerCreater) CreateCommandsHandlers(noteBackClient *external.NoteBackendClient, authServerClient *external.AuthServerClient, d dao.TelegramBotDao) []bot.Option {
	var result = []bot.Option{}
	for i, value := range h.Config.Server.HandlersToInit {
		log.Println(fmt.Sprintf("CreateHandlers handler : %s. With order: %x", value, i+1))
		switch value {
		case START_HANDLER:
			{
				result = append(result, command.NewStartCommandHandler(d).Init()...)
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
				result = append(result, command.NewNoteCommandHandler(noteBackClient).Init()...)
			}

		case PROFILE_HANDLER:
			{
				result = append(result, command.NewProfileCommandHandler(authServerClient).Init()...)
			}
		}
	}
	return result
}
