package usecases

import (
	apiDto "authServer/api/docs"
	keycloak "authServer/external"
	authDaoInterface "authServer/internal/dao/interface"
	domain "authServer/internal/domain"
)

const (
	TG  apiDto.CreateAccountRequestAccountType = "TG"
	WEB apiDto.CreateAccountRequestAccountType = "WEB"
)

type RegisterAccountUseCase struct {
	Keycloak *keycloak.KeycloakClient
	Dao      authDaoInterface.AuthDao
}

func (register *RegisterAccountUseCase) Register(request apiDto.CreateAccountRequest) (result domain.Account, err error) {
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
			FirstName: request.FirstName,
			LastName:  request.LastName,
			Email:     request.Email,
			Login:     *request.Login,
			Type:      string(*request.AccountType),
		},
	), nil
}
