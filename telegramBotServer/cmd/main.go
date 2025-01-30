package main

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"os/signal"
	"telegramCLient/config"
	"telegramCLient/internal/creater"
)

// ______________________________________________________________________
func main() {
	appConfig := loadConfig("MODE")

	ctx, cancel := getContext()
	defer cancel()

	opts := creater.CreateHandlerStarter(appConfig, make(map[int64][]int)).CreateHandlers()
	createAndStartBot(ctx, appConfig, opts)
}

//______________________________________________________________________

func createAndStartBot(ctx context.Context, appConfig *config.AppConfig, opts []bot.Option) {
	b, err := bot.New(appConfig.BotToken, opts...)
	if nil != err {
		panic(err)
	}
	b.Start(ctx)
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
