package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"telegramCLient/config"
	"telegramCLient/internal/bot"
	"telegramCLient/internal/handler"
)

// ______________________________________________________________________
func main() {
	appConfig := loadConfig("MODE")
	tempTgMessageId := make(map[int64][]int64)
	bot, dispatcher := bot.NewNoteTGBot(appConfig).Init()
	handler.CreateHandlerStarter(appConfig, dispatcher, tempTgMessageId).InitAndStart()
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
