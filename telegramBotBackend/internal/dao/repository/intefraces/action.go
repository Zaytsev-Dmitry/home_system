package intefraces

import "telegramCLient/internal/domain"

type ActionRepository interface {
	SaveOrUpdate(telegramUserId int64, userInput bool, lastSentMessageId int, commandName string, isRunning bool)
	Update(telegramUserId int64, needUserAction bool, lastSentMessageId int, isRunning bool)
	GetByTgId(telegramUserId int64) domain.UserAction
	CloseConnection()
}
