package main

import (
	"botCoreService/internal/app/service"
	"botCoreService/internal/infrastructure/transport/tg/commands"
	"botCoreService/internal/infrastructure/transport/tg/commands/registry"
	"botCoreService/internal/infrastructure/transport/tg/commands/start"
	"botCoreService/pkg/config_loader"
	"botCoreService/pkg/telegram"
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"log"
	"os"
	"os/signal"
)

const (
	START_COMMAND string = "START"
)

func main() {
	//гружу конфиг
	appConfig := config_loader.LoadConfig()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	bot, err := bot.New(appConfig.Telegram.BotToken)
	if nil != err {
		panic(err)
	}

	createAndRegisterCommands(appConfig, bot, ctx)
	bot.Start(ctx)
}

func createAndRegisterCommands(conf *config_loader.AppConfig, b *bot.Bot, ctx context.Context) {
	for i, value := range conf.Telegram.CommandsToInit {
		var newCommand commands.BaseCommand
		log.Println(fmt.Sprintf("Create command : %s. With order: %x", value, i+1))

		commandRegistry := registry.NewCommandSessionManager()
		sessionManager := service.NewInMemorySessionStorage()
		switch value {
		case START_COMMAND:
			{
				newCommand = start.NewStartCommand(b, ctx, commandRegistry, sessionManager)
			}
		default:
			{
				fmt.Println("Неизвестная команда")
			}
		}
		newCommand.RegisterHandler()

		//хэндлер обработки ответов (вводов) от пользователя
		b.RegisterHandler(bot.HandlerTypeMessageText, "", bot.MatchTypeContains, func(ctx context.Context, b *bot.Bot, update *models.Update) {
			userId := telegram.GetUserId(update)
			if cmd, ok := commandRegistry.Get(userId); ok {
				cmd.ProceedUserInput(ctx, b, update)
			} else {
				b.SendMessage(ctx, &bot.SendMessageParams{
					ChatID: userId,
					Text:   "Не понимаю к чему относится твое сообщение :(",
				})
			}
		})
	}
}
