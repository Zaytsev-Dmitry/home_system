package authServerPresenter

import (
	domain "authServer/internal/domain"
	authSpec "github.com/Zaytsev-Dmitry/home_system_open_api/authServerBackend"
)

type Presenter struct {
}

func (presenter *Presenter) ToAccountResponse(entity domain.Account) authSpec.AccountResponse {
	return authSpec.AccountResponse{
		Email:     entity.Email,
		FirstName: entity.FirstName,
		LastName:  entity.LastName,
		Login:     &entity.Login,
	}
}

func (presenter *Presenter) ToEntity(requestEntity *authSpec.CreateAccountRequest) domain.Account {
	return domain.Account{
		FirstName: requestEntity.FirstName,
		LastName:  requestEntity.LastName,
		Login:     *requestEntity.Login,
		Email:     requestEntity.Email,
	}
}
