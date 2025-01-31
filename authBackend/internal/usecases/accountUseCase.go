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

type RegisterAccountUseCase struct {
	Keycloak       *keycloak.KeycloakClient
	Repo           account.AccountRepository
	ProfileUsecase ProfileUseCase
}

func (register *RegisterAccountUseCase) Register(request authSpec.CreateAccountRequest) (result domain.Account, err error) {
	switch *request.AccountType {
	case WEB:
		{
			errKeycloak := register.Keycloak.RegisterAccount(request)
			if errKeycloak != nil {
				return domain.Account{}, errKeycloak
			}
		}
	}
	saved := register.Repo.Save(
		domain.Account{
			FirstName:  request.FirstName,
			LastName:   request.LastName,
			Email:      *request.Email,
			Type:       string(*request.AccountType),
			TelegramId: *request.TelegramId,
			IsActive:   true,
		},
	)
	register.ProfileUsecase.Create(saved, *request.TelegramUsername)
	return saved, nil
}
