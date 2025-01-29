package main

import (
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"telegramCLient/config"
	"telegramCLient/internal/bot"
	"telegramCLient/internal/handler"
	thCommandHandlerinterface "telegramCLient/internal/handler/interface"
)

const (
	START_HANDLER    string = "StartCommandHandler"
	REGISTER_HANDLER string = "RegisterCommandHandler"
)

func createAndInitHandlers(config *config.AppConfig, disp *ext.Dispatcher) {
	for i, value := range config.HandlersToInit {
		log.Println(fmt.Sprintf("Create handler : %s. With order: %x", value, i+1))
		var createdHandler thCommandHandlerinterface.TGCommandHandler
		switch value {
		case START_HANDLER:
			{
				createdHandler = &handler.StartCommandHandler{}
			}
		case REGISTER_HANDLER:
			{
				createdHandler = &handler.RegisterUserCommandHandler{}
			}
		}
		createdHandler.Init(disp)
	}
}

// ______________________________________________________________________
func main() {
	appConfig := loadConfig("MODE")
	bot, dispatcher := bot.NewNoteTGBot(appConfig).Init()
	createAndInitHandlers(appConfig, dispatcher)
	bot.Idle()
}

//______________________________________________________________________

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
