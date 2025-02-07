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

type AccountUseCase struct {
	Keycloak *keycloak.KeycloakClient
	Repo     intefraces.AccountRepository
}

func (usecase *AccountUseCase) Register(request authSpec.CreateAccountRequest) (domain.Account, error, int) {
	var status = http.StatusOK
	var result domain.Account

	result, respErr := usecase.runBusinessLayout(request)
	status = usecase.getStatus(respErr, status)
	return result, respErr, status
}

func (usecase *AccountUseCase) getStatus(respErr error, status int) int {
	if respErr != nil {
		utilities.GetLogger().Error(respErr.Error())
		status = http.StatusInternalServerError
	}
	return status
}

func (usecase *AccountUseCase) runBusinessLayout(request authSpec.CreateAccountRequest) (domain.Account, error) {
	var result domain.Account
	respErr, keycloakEntity := usecase.getKeycloakUser(request)
	if respErr == nil {
		result, respErr = usecase.Repo.CreateAccountAndProfile(keycloakEntity, *request.TelegramId)
	}
	return result, respErr
}

func (usecase *AccountUseCase) getKeycloakUser(request authSpec.CreateAccountRequest) (error, keycloak.KeycloakEntity) {
	var keycloakEntity keycloak.KeycloakEntity
	var err error

	err, keycloakEntity = usecase.Keycloak.RegisterAccount(request)
	if err != nil {
		if errors.Is(err, keycloak.Conflict409) {
			//пользак уже есть в keycloak и соответственно в базе
			utilities.GetLogger().Warn(err.Error())
			keycloakEntity, err = usecase.Keycloak.GetUser(*request.Email)
			if err != nil {
				utilities.GetLogger().Error(err.Error())
			}
		} else {
			utilities.GetLogger().Error(err.Error())
		}
	} else {
		keycloakEntity, err = usecase.Keycloak.GetUser(*request.Email)
	}
	return err, keycloakEntity
}

func (usecase *AccountUseCase) GetAccountIdByTgId(tgId int64) (accId int64) {
	return usecase.Repo.GetIdByTgId(tgId)
}
