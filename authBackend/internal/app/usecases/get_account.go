package usecases

import "authServer/internal/app/domain"

type GetAccByTelegramIdUCase interface {
	Get(telegramId int64) (domain.Account, error)
	GetAccountIdByTgId(tgId int64) (accId int64)
}
