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
	Keycloak       *keycloak.KeycloakClient
	Repo           intefraces.AccountRepository
	ProfileUsecase ProfileUseCase
}

func (usecase *AccountUseCase) Register(request authSpec.CreateAccountRequest) (domain.Account, error, int) {
	var status = http.StatusOK
	var respErr error
	var keycloakEntity keycloak.KeycloakEntity
	var result domain.Account

	respErr, result = usecase.runBusinessLayout(request, respErr, keycloakEntity, result)
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

func (usecase *AccountUseCase) runBusinessLayout(request authSpec.CreateAccountRequest, respErr error, keycloakEntity keycloak.KeycloakEntity, result domain.Account) (error, domain.Account) {
	respErr, keycloakEntity = usecase.getKeycloakUser(request)
	if respErr == nil {
		result, respErr = usecase.getDaoEntity(request, keycloakEntity)
	}
	if respErr == nil {
		usecase.ProfileUsecase.Create(result, *request.TelegramUsername)
	}
	return respErr, result
}

func (usecase *AccountUseCase) getDaoEntity(request authSpec.CreateAccountRequest, keycloakEntity keycloak.KeycloakEntity) (domain.Account, error) {
	saved, respErr := usecase.Repo.Register(
		domain.Account{
			FirstName:  &keycloakEntity.FirstName,
			LastName:   &keycloakEntity.LastName,
			Email:      keycloakEntity.Email,
			TelegramId: request.TelegramId,
			IsActive:   true,
		},
	)
	return saved, respErr
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
	}
	return err, keycloakEntity
}

func (usecase *AccountUseCase) GetAccountIdByTgId(tgId int64) (accId int64) {
	return usecase.Repo.GetIdByTgId(tgId)
}
