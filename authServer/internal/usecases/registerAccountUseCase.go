package usecases

import (
	apiDto "authServer/api/docs"
	kecyloak "authServer/external"
	domain "authServer/internal/domain"
)

type RegisterAccountUseCase struct {
}

var keycloakClient = kecyloak.KeycloakClient{}

func (register *RegisterAccountUseCase) Register(request apiDto.CreateAccountRequest) (result domain.AccountEntity, err error) {
	keycloakClient.RegisterAccount(request)
	return domain.AccountEntity{}, nil
}
