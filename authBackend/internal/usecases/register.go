package usecases

import (
	generatedApi "authServer/api/http"
	"authServer/external/keycloak"
	"authServer/internal/dao/repository/intefraces"
	domain "authServer/internal/domain"
	"authServer/pkg/utilities"
	"errors"
)

type RegisterAccount struct {
	Keycloak *keycloak.KeycloakClient
	Repo     intefraces.AccountRepository
}

func (usecase *RegisterAccount) Register(request generatedApi.CreateAccountRequest) (domain.Account, error) {
	var result domain.Account
	result, respErr := usecase.runBusinessLayout(request)
	return result, respErr
}

func (usecase *RegisterAccount) runBusinessLayout(request generatedApi.CreateAccountRequest) (domain.Account, error) {
	var result domain.Account
	err, keycloakEntity := usecase.getKeycloakUser(request)
	if err == nil {
		result, err = usecase.Repo.CreateAccountAndProfile(keycloakEntity, *request.Username, *request.TelegramId)
	}
	return result, err
}

func (usecase *RegisterAccount) getKeycloakUser(request generatedApi.CreateAccountRequest) (error, keycloak.KeycloakEntity) {
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
