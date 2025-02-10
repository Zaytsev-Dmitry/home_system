package intefraces

import "telegramCLient/internal/domain"

type ActionRepository interface {
	Save(telegramUserId int64, lastAction string, lastRequirement string, lastSentMessageId int)
	Update(telegramUserId int64, lastAction string, lastRequirement string, lastSentMessageId int)
	GetByTgId(telegramUserId int64) domain.UserAction
	CloseConnection()
}
