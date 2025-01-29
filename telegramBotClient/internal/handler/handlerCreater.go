package handler

import (
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"log"
	"telegramCLient/config"
	"telegramCLient/internal/handler/interface"
)

const (
	START_HANDLER    string = "StartCommandHandler"
	REGISTER_HANDLER string = "RegisterCommandHandler"
)

type HandlerStarter struct {
	Config                *config.AppConfig
	Disp                  *ext.Dispatcher
	tempMessageCollection map[int64][]int64
}

func CreateHandlerStarter(conf *config.AppConfig, disp *ext.Dispatcher, tempMessage map[int64][]int64) *HandlerStarter {
	return &HandlerStarter{
		Config:                conf,
		Disp:                  disp,
		tempMessageCollection: tempMessage,
	}
}

// TODO отловаить ошибки
func (h *HandlerStarter) InitAndStart() {
	for i, value := range h.Config.HandlersToInit {
		log.Println(fmt.Sprintf("Create handler : %s. With order: %x", value, i+1))
		var createdHandler thCommandHandlerinterface.TGCommandHandler
		switch value {
		case START_HANDLER:
			{
				createdHandler = NewStartCommandHandler(h.tempMessageCollection)
			}
		case REGISTER_HANDLER:
			{
				createdHandler = NewRegisterUserCommandHandler(h.Config.AuthServerUrl, h.tempMessageCollection)
			}
		}
		createdHandler.Init(h.Disp)
	}
}
