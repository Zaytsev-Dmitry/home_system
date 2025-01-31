package usecases

import (
	keycloak "authServer/external"
	"authServer/internal/dao/repository/account"
	domain "authServer/internal/domain"
	authSpec "github.com/Zaytsev-Dmitry/home_system_open_api/authServerBackend"
)

const (
	TG  authSpec.CreateAccountRequestAccountType = "TG"
	WEB authSpec.CreateAccountRequestAccountType = "WEB"
)

type AccountUseCase struct {
	Keycloak       *keycloak.KeycloakClient
	Repo           account.AccountRepository
	ProfileUsecase ProfileUseCase
}

func (usecase *AccountUseCase) Register(request authSpec.CreateAccountRequest) (result domain.Account, err error) {
	switch *request.AccountType {
	case WEB:
		{
			errKeycloak := usecase.Keycloak.RegisterAccount(request)
			if errKeycloak != nil {
				return domain.Account{}, errKeycloak
			}
		}
	}
	saved := usecase.Repo.Save(
		domain.Account{
			FirstName:  request.FirstName,
			LastName:   request.LastName,
			Email:      *request.Email,
			Type:       string(*request.AccountType),
			TelegramId: *request.TelegramId,
			IsActive:   true,
		},
	)
	usecase.ProfileUsecase.Create(saved, *request.TelegramUsername)
	return saved, nil
}

func (usecase *AccountUseCase) GetAccountIdByTgId(tgId int64) (accId int64) {
	return usecase.Repo.GetIdByTgId(tgId)
}
