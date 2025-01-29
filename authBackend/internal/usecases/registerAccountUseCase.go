package usecases

import (
	keycloak "authServer/external"
	authDaoInterface "authServer/internal/dao/interface"
	domain "authServer/internal/domain"
	authSpec "github.com/Zaytsev-Dmitry/home_system_open_api/authServerBackend"
)

const (
	TG  authSpec.CreateAccountRequestAccountType = "TG"
	WEB authSpec.CreateAccountRequestAccountType = "WEB"
)

type RegisterAccountUseCase struct {
	Keycloak *keycloak.KeycloakClient
	Dao      authDaoInterface.AuthDao
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
	return register.Dao.Save(
		domain.Account{
			FirstName:  request.FirstName,
			LastName:   request.LastName,
			Email:      request.Email,
			Login:      *request.Login,
			Type:       string(*request.AccountType),
			TelegramId: request.TelegramId,
		},
	), nil
}
