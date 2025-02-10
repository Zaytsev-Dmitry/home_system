package main

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	_ "github.com/lib/pq"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"os/signal"
	"telegramCLient/config"
	"telegramCLient/external"
	"telegramCLient/internal/action"
	"telegramCLient/internal/creater"
	"telegramCLient/internal/dao"
	"telegramCLient/internal/handler/command"
)

// ______________________________________________________________________
func main() {
	appConfig := loadConfig("MODE")
	ctx, cancel := getContext()
	defer cancel()

	dao := dao.CreateDao(*appConfig)

	opts, commands := creater.CreateHandlerStarter(appConfig).CreateCommandsHandlers(
		external.NewNoteBackendClient(appConfig.Server.NoteBackendUrl),
		external.NewAuthServerClient(appConfig.Server.AuthServerUrl),
		*dao,
	)
	createAndStartBot(ctx, appConfig, opts, *dao, commands)
	defer dao.Close()
}

//______________________________________________________________________

func createAndStartBot(ctx context.Context, appConfig *config.AppConfig, opts []bot.Option, d dao.TelegramBotDao, commands []command.BaseCommand) {
	b, err := bot.New(appConfig.Server.BotToken, opts...)
	if nil != err {
		panic(err)
	}
	userAction := action.NewAction(d, commands)
	b.RegisterHandler(bot.HandlerTypeMessageText, "", bot.MatchTypeContains, func(ctx context.Context, b *bot.Bot, update *models.Update) {
		userAction.Proceed(ctx, b, update)
	})
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
