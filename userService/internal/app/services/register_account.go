package services

import (
	"errors"
	generatedApi "userService/api/http"
	"userService/internal/app/domain"
	"userService/internal/app/ports/out/dao/repository/account"
	"userService/internal/app/ports/out/keycloak"
	"userService/pkg/utilities"
)

type RegisterAccountUseCaseImpl struct {
	Keycloak *keycloak.KeycloakClient
	Repo     account.AccountRepository
}

func (u *RegisterAccountUseCaseImpl) Register(request generatedApi.CreateAccountRequest) (domain.Account, error) {
	var result domain.Account
	result, respErr := u.runBusinessLayout(request)
	return result, respErr
}

func (u *RegisterAccountUseCaseImpl) runBusinessLayout(request generatedApi.CreateAccountRequest) (domain.Account, error) {
	var result domain.Account
	err, keycloakEntity := u.getKeycloakUser(request)
	if err == nil {
		result, err = u.Repo.CreateAccountAndProfile(keycloakEntity, *request.Username, *request.TelegramId)
	}
	return result, err
}

func (u *RegisterAccountUseCaseImpl) getKeycloakUser(request generatedApi.CreateAccountRequest) (error, keycloak.KeycloakEntity) {
	err, keycloakEntity := u.Keycloak.RegisterAccount(request)
	if err != nil {
		if errors.Is(err, keycloak.Conflict409) {
			//пользак уже есть в keycloak и соответственно в базе
			keycloakEntity, err = u.Keycloak.GetUser(*request.Email)
			if err != nil {
				utilities.GetLogger().Error(err.Error())
			}
		} else {
			utilities.GetLogger().Error(err.Error())
		}
	} else {
		keycloakEntity, err = u.Keycloak.GetUser(*request.Email)
	}
	return err, keycloakEntity
}
