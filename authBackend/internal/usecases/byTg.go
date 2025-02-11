package usecases

import (
	"authServer/internal/dao/repository/intefraces"
	domain "authServer/internal/domain"
	"net/http"
)

type GetAccByTelegramId struct {
	Repo intefraces.AccountRepository
}

func (controller *GetAccByTelegramId) Get(telegramId int64) (domain.Account, error, int) {
	entity, err := controller.Repo.GetByTgId(telegramId)
	return entity, err, IfExistErrLogAndReturn500Http(err, http.StatusOK)
}

func (usecase *GetAccByTelegramId) GetAccountIdByTgId(tgId int64) (accId int64) {
	return usecase.Repo.GetIdByTgId(tgId)
}
