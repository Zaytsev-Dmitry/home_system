package creater

import (
	"telegramCLient/config"
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
//func (h *HandlerCreater) CreateCommandsHandlers(noteBackClient *external.NoteBackendClient, authServerClient *external.AuthServerClient, d dao.TelegramBotDao) ([]bot.Option, []command.BaseCommand) {
//	var result = []bot.Option{}
//	var commands = []command.BaseCommand{}
//	for i, value := range h.Config.Server.HandlersToInit {
//		var newCommand command.BaseCommand
//		log.Println(fmt.Sprintf("CreateHandlers handler : %s. With order: %x", value, i+1))
//		switch value {
//
//		case START_HANDLER:
//			{
//				newCommand = command.NewStartCommandHandler(d, *authServerClient)
//			}
//		case TUTORIAL_HANDLER:
//			{
//				newCommand = command.NewTutorialCommandHandler()
//			}
//		case MENU_HANDLER:
//			{
//				newCommand = menu.NewMenuCommandHandler()
//			}
//		case NOTE_HANDLER:
//			{
//				newCommand = command.NewNoteCommandHandler(d, noteBackClient)
//			}
//
//		case PROFILE_HANDLER:
//			{
//				newCommand = command.NewProfileCommandHandler(authServerClient)
//			}
//		}
//		result = append(result, newCommand.Init()...)
//		commands = append(commands, newCommand)
//	}
//	return result, commands
//}
