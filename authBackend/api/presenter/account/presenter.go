package account

import (
	domain "authServer/internal/domain"
	authSpec "github.com/Zaytsev-Dmitry/home_system_open_api/authServerBackend"
)

type Presenter struct {
}

func (presenter *Presenter) ToAccountResponse(entity domain.Account) authSpec.AccountResponse {
	return authSpec.AccountResponse{
		Email:     &entity.Email,
		FirstName: entity.FirstName,
		LastName:  entity.LastName,
	}
}

func (presenter *Presenter) ToEntity(requestEntity *authSpec.CreateAccountRequest) domain.Account {
	return domain.Account{
		FirstName: requestEntity.FirstName,
		LastName:  requestEntity.LastName,
		Email:     *requestEntity.Email,
	}
}
