package registry

import (
	"botCoreService/internal/infrastructure/transport/tg/commands"
)

type CommandRegistry interface {
	Set(userID int64, cmd commands.BaseCommand)
	Get(userID int64) (commands.BaseCommand, bool)
	Delete(userID int64)
}
