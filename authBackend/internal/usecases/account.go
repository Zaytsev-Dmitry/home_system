package usecases

import (
	"authServer/external/keycloak"
	"authServer/internal/dao/repository/intefraces"
	domain "authServer/internal/domain"
	"authServer/pkg/utilities"
	"errors"
	authSpec "github.com/Zaytsev-Dmitry/home_system_open_api/authServerBackend"
	"net/http"
)

const (
	TG  authSpec.CreateAccountRequestAccountType = "TG"
	WEB authSpec.CreateAccountRequestAccountType = "WEB"
)

type AccountUseCase struct {
	Keycloak       *keycloak.KeycloakClient
	Repo           intefraces.AccountRepository
	ProfileUsecase ProfileUseCase
}

func (usecase *AccountUseCase) Register(request authSpec.CreateAccountRequest) (domain.Account, error, int) {
	var status int
	if *request.AccountType == WEB {
		err, account := usecase.Keycloak.RegisterAccount(request)
		if err != nil {
			if errors.Is(err, keycloak.Conflict409) {
				utilities.GetLogger().Warn(err.Error())
				//todo select из базы
				return account, nil, http.StatusOK
			} else {
				utilities.GetLogger().Error(err.Error())
				status = http.StatusInternalServerError
			}
			return account, err, status
		}
	}
	saved, err := usecase.Repo.Save(
		domain.Account{
			FirstName:  request.FirstName,
			LastName:   request.LastName,
			Email:      *request.Email,
			Type:       string(*request.AccountType),
			TelegramId: request.TelegramId,
			IsActive:   true,
		},
	)
	if err != nil {
		status = http.StatusInternalServerError
	}
	usecase.ProfileUsecase.Create(saved, *request.TelegramUsername)
	return saved, nil, http.StatusOK
}

func (usecase *AccountUseCase) GetAccountIdByTgId(tgId int64) (accId int64) {
	return usecase.Repo.GetIdByTgId(tgId)
}
