package main

import (
	"fmt"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"telegramCLient/config"
)

func main() {
	bot := getBot()
	updates, _ := bot.UpdatesViaLongPolling(nil)
	bh, _ := th.NewBotHandler(bot, updates)
	defer bh.Stop()
	defer bot.StopLongPolling()

	bh.HandleMessage(func(bot *telego.Bot, message telego.Message) {
		_, _ = bot.SendMessage(tu.Messagef(
			tu.ID(message.Chat.ID),
			"Привет %s!", message.From.Username+" "+"сначала тебе необходимо зарегистрироваться",
		).WithReplyMarkup(tu.InlineKeyboard(
			tu.InlineKeyboardRow(tu.InlineKeyboardButton("Go!").WithCallbackData("register"))),
		))
	}, th.CommandEqual("start"))

	bh.HandleCallbackQuery(func(bot *telego.Bot, query telego.CallbackQuery) {
		_, _ = bot.SendMessage(tu.Message(tu.ID(query.Message.GetChat().ID), "GO"))
		_ = bot.AnswerCallbackQuery(tu.CallbackQuery(query.ID).WithText("Done"))
	}, th.AnyCallbackQueryWithMessage(), th.CallbackDataEqual("register"))

	bh.Start()
}

func getBot() *telego.Bot {
	appConfig := loadConfig("MODE")
	bot, err := telego.NewBot(appConfig.BotToken, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return bot
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
