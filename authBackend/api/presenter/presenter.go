package authServerPresenter

import (
	apiDTO "authServer/api/docs"
	domain "authServer/internal/domain"
)

type Presenter struct {
}

func (presenter *Presenter) ToAccountResponse(entity domain.Account) apiDTO.AccountResponse {
	return apiDTO.AccountResponse{
		Email:     entity.Email,
		FirstName: entity.FirstName,
		LastName:  entity.LastName,
		Login:     &entity.Login,
	}
}

func (presenter *Presenter) ToEntity(requestEntity *apiDTO.CreateAccountRequest) domain.Account {
	return domain.Account{
		FirstName: requestEntity.FirstName,
		LastName:  requestEntity.LastName,
		Login:     *requestEntity.Login,
		Email:     requestEntity.Email,
	}
}
