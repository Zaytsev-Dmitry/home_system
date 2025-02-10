package intefraces

import "telegramCLient/internal/domain"

type ActionRepository interface {
	SaveOrUpdate(telegramUserId int64, commandState string, needUserAction bool, lastSentMessageId int, commandName string)
	Update(telegramUserId int64, commandState string, needUserAction bool, lastSentMessageId int)
	GetByTgId(telegramUserId int64) domain.UserAction
	CloseConnection()
}
