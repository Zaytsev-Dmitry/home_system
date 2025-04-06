package services

import (
	"authServer/internal/app/domain"
	"authServer/internal/app/ports/out/dao/repository/intefraces"
)

type GetAccountUCaseImpl struct {
	Repo intefraces.AccountRepository
}

func (controller *GetAccountUCaseImpl) Get(telegramId int64) (domain.Account, error) {
	entity, err := controller.Repo.GetByTgId(telegramId)
	return entity, err
}

func (usecase *GetAccountUCaseImpl) GetAccountIdByTgId(tgId int64) (accId int64) {
	return usecase.Repo.GetIdByTgId(tgId)
}
