package presenter

import (
	generatedApi "authServer/api/http"
	domain "authServer/internal/domain"
)

type AccountPresenter struct {
}

func (presenter *AccountPresenter) ToAccountResponse(entity domain.Account) generatedApi.AccountResponse {
	return generatedApi.AccountResponse{
		Email:            &entity.Email,
		FirstName:        &entity.FirstName,
		Id:               &entity.ID,
		LastName:         &entity.LastName,
		TelegramId:       &entity.TelegramId,
		TelegramUserName: &entity.TelegramUserName,
		Username:         &entity.Username,
	}
}

func (presenter *AccountPresenter) ToEntity(requestEntity *generatedApi.CreateAccountRequest) domain.Account {
	return domain.Account{
		FirstName: *requestEntity.FirstName,
		LastName:  *requestEntity.LastName,
		Email:     *requestEntity.Email,
	}
}
