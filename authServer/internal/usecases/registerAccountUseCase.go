package usecases

import (
	apiDto "authServer/api/docs"
	keycloak "authServer/external"
	domain "authServer/internal/domain"
)

type RegisterAccountUseCase struct {
	Keycloak *keycloak.KeycloakClient
}

func (register *RegisterAccountUseCase) Register(request apiDto.CreateAccountRequest) (result domain.AccountEntity, err error) {
	register.Keycloak.RegisterAccount(request)
	return domain.AccountEntity{}, nil
}
