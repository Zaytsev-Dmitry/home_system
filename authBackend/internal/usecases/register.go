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

type RegisterAccount struct {
	Keycloak *keycloak.KeycloakClient
	Repo     intefraces.AccountRepository
}

func (usecase *RegisterAccount) Register(request authSpec.CreateAccountRequest) (domain.Account, error, int) {
	var status = http.StatusOK
	var result domain.Account
	result, respErr := usecase.runBusinessLayout(request)
	return result, respErr, IfExistErrLogAndReturn500Http(respErr, status)
}

func (usecase *RegisterAccount) runBusinessLayout(request authSpec.CreateAccountRequest) (domain.Account, error) {
	var result domain.Account
	respErr, keycloakEntity := usecase.getKeycloakUser(request)
	if respErr == nil {
		result, respErr = usecase.Repo.CreateAccountAndProfile(keycloakEntity, *request.Username, *request.TelegramId)
	}
	return result, respErr
}

func (usecase *RegisterAccount) getKeycloakUser(request authSpec.CreateAccountRequest) (error, keycloak.KeycloakEntity) {
	err, keycloakEntity := usecase.Keycloak.RegisterAccount(request)
	if err != nil {
		if errors.Is(err, keycloak.Conflict409) {
			//пользак уже есть в keycloak и соответственно в базе
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
