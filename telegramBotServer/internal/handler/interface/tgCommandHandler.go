package thCommandHandlerinterface

import "github.com/PaulSonOfLars/gotgbot/v2/ext"

type TGCommandHandler interface {
	Init(dispatcher *ext.Dispatcher)
}
