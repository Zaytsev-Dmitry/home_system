package usecases

import (
	"authServer/internal/dao/repository/intefraces"
	domain "authServer/internal/domain"
)

type GetAccByTelegramId struct {
	Repo intefraces.AccountRepository
}

func (controller *GetAccByTelegramId) Get(telegramId int64) (domain.Account, error) {
	entity, err := controller.Repo.GetByTgId(telegramId)
	return entity, err
}

func (usecase *GetAccByTelegramId) GetAccountIdByTgId(tgId int64) (accId int64) {
	return usecase.Repo.GetIdByTgId(tgId)
}
