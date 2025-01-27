package main

import (
	"fmt"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"telegramCLient/config"
	"telegramCLient/internal/handler"
)

func main() {
	bh, bot := initNoteTGBot()
	defer bh.Stop()
	defer bot.StopLongPolling()

	handler.HandleStartCommand(bh)
	bh.Start()
}

func initNoteTGBot() (bh *th.BotHandler, bot *telego.Bot) {
	appConfig := loadConfig("MODE")
	bot, err := telego.NewBot(appConfig.BotToken, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	updates, _ := bot.UpdatesViaLongPolling(nil)
	bh, _ = th.NewBotHandler(bot, updates)
	return bh, bot
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
