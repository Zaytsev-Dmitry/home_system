package usecases

import "authBackend/internal/app/domain"

type GetAccountUCase interface {
	Get(telegramId int64) (domain.Account, error)
	GetAccountIdByTgId(tgId int64) (accId int64)
}
