package main

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/go-telegram/bot"
	_ "github.com/lib/pq"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"os/signal"
	"telegramCLient/config"
	"telegramCLient/external"
	"telegramCLient/internal/dao"
	"telegramCLient/internal/handler/command"
	"telegramCLient/internal/handler/command/menu"
	"telegramCLient/internal/handler/command/notes"
	"telegramCLient/internal/handler/command/profile"
	"telegramCLient/internal/handler/command/start"
	"telegramCLient/internal/handler/command/test"
	"telegramCLient/internal/handler/command/tutorial"
	"telegramCLient/internal/storage"
)

const (
	TEST_COMMAND     string = "TEST"
	MENU_COMMAND     string = "MENU"
	TUTORIAL_COMMAND string = "TUTORIAL"
	PROFILE_COMMAND  string = "PROFILE"
	NOTES_COMMAND    string = "NOTES"
	START_COMMAND    string = "START"
)

// ______________________________________________________________________
func main() {
	appConfig := loadConfig("MODE")
	ctx, cancel := getContext()
	defer cancel()
	dao := dao.CreateDao(*appConfig)

	bot, err := bot.New(appConfig.Server.BotToken)
	if nil != err {
		panic(err)
	}

	createAndRegisterCommands(appConfig, bot, ctx, *dao)
	bot.Start(ctx)
	defer dao.Close()
}

func createAndRegisterCommands(conf *config.AppConfig, bot *bot.Bot, ctx context.Context, dao dao.TelegramBotDao) {
	storage := *storage.NewStorage()
	authServerClient := external.NewAuthServerClient(conf.Server.AuthServerUrl)
	noteBackendClient := external.NewNoteBackendClient(conf.Server.NoteBackendUrl)
	for i, value := range conf.Server.CommandsToInit {
		var newCommand command.BaseCommand
		log.Println(fmt.Sprintf("Create command : %s. With order: %x", value, i+1))
		switch value {
		case TEST_COMMAND:
			{
				newCommand = test.NewTestCommand(storage, bot, ctx)
			}
		case MENU_COMMAND:
			newCommand = menu.NewMenuCommand(storage, bot, ctx)

		case TUTORIAL_COMMAND:
			newCommand = tutorial.NewTutorialCommand(storage, bot, ctx)

		case PROFILE_COMMAND:
			newCommand = profile.NewProfileCommand(storage, bot, ctx, authServerClient)

		case NOTES_COMMAND:
			newCommand = notes.NewNotesCommand(storage, bot, ctx, dao, noteBackendClient)
		case START_COMMAND:
			newCommand = start.NewStartCommand(storage, bot, ctx, dao, authServerClient)
		default:
			fmt.Println("Неизвестная команда")
		}
		newCommand.RegisterHandler()
	}
}

func getContext() (ctx context.Context, stop context.CancelFunc) {
	return signal.NotifyContext(context.Background(), os.Interrupt)
}

func loadConfig(env string) *config.AppConfig {
	var appProfile = "config/" + "%s" + ".yaml"
	getenv := os.Getenv(env)
	switch getenv {
	case "dev":
		appProfile = fmt.Sprintf(appProfile, "dev")
	case "test":
		appProfile = fmt.Sprintf(appProfile, "test")
	}
	log.Println(fmt.Sprintf("Run application in mode : %s", getenv))
	f, err := os.Open(appProfile)
	if err != nil {
	}
	defer f.Close()

	var cfg config.AppConfig
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {

	}
	return &cfg
}
