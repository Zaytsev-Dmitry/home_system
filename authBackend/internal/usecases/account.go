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
	var respErr error
	if *request.AccountType == WEB {
		err, account := usecase.Keycloak.RegisterAccount(request)
		if err != nil {
			if errors.Is(err, keycloak.Conflict409) {
				//пользак уже есть в keycloak и соответственно в базе
				utilities.GetLogger().Warn(err.Error())
				//TODO сходить в keycloak
				return account, nil, http.StatusOK
			} else {
				utilities.GetLogger().Error(err.Error())
				status = http.StatusInternalServerError
			}
			return account, err, status
		}
	}
	saved, regUserErr := usecase.Repo.Register(
		domain.Account{
			FirstName:  request.FirstName,
			LastName:   request.LastName,
			Email:      *request.Email,
			Type:       string(*request.AccountType),
			TelegramId: request.TelegramId,
			IsActive:   true,
		},
	)
	if regUserErr != nil {
		status = http.StatusInternalServerError
	}
	usecase.ProfileUsecase.Create(saved, *request.TelegramUsername)
	return saved, respErr, status
}

func (usecase *AccountUseCase) GetAccountIdByTgId(tgId int64) (accId int64) {
	return usecase.Repo.GetIdByTgId(tgId)
}
