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
	var keycloakEntity keycloak.KeycloakEntity

	if *request.AccountType == WEB {
		respErr, keycloakEntity = usecase.Keycloak.RegisterAccount(request)
		if respErr != nil {
			if errors.Is(respErr, keycloak.Conflict409) {
				//пользак уже есть в keycloak и соответственно в базе
				utilities.GetLogger().Warn(respErr.Error())
				user, getUserErr := usecase.Keycloak.GetUser(*request.Email)
				keycloakEntity = user
				if getUserErr != nil {
					utilities.GetLogger().Error(respErr.Error())
					status = http.StatusInternalServerError
				}
			} else {
				utilities.GetLogger().Error(respErr.Error())
				status = http.StatusInternalServerError
			}
		}
	}
	saved, respErr := usecase.Repo.Register(
		domain.Account{
			FirstName:  &keycloakEntity.FirstName,
			LastName:   &keycloakEntity.LastName,
			Email:      keycloakEntity.Email,
			Type:       string(*request.AccountType),
			TelegramId: request.TelegramId,
			IsActive:   true,
		},
	)
	if respErr != nil {
		utilities.GetLogger().Error(respErr.Error())
		status = http.StatusInternalServerError
	} else {
		usecase.ProfileUsecase.Create(saved, *request.TelegramUsername)
	}
	return saved, respErr, status
}

func (usecase *AccountUseCase) GetAccountIdByTgId(tgId int64) (accId int64) {
	return usecase.Repo.GetIdByTgId(tgId)
}
