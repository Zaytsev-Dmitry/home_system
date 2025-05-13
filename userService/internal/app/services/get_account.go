package services

import (
	"userService/internal/app/domain"
	"userService/internal/app/ports/out/dao/repository/account"
)

type GetAccountUCaseImpl struct {
	Repo account.AccountRepository
}

func (controller *GetAccountUCaseImpl) Get(telegramId int64) (domain.Account, error) {
	entity, err := controller.Repo.GetByTgId(telegramId)
	return entity, err
}

func (usecase *GetAccountUCaseImpl) GetAccountIdByTgId(tgId int64) (accId int64) {
	return usecase.Repo.GetIdByTgId(tgId)
}
