package presenter

import (
	domain "authServer/internal/domain"
	authSpec "github.com/Zaytsev-Dmitry/home_system_open_api/authServerBackend"
)

type AccountPresenter struct {
}

func (presenter *AccountPresenter) ToAccountResponse(entity domain.Account) authSpec.AccountResponse {
	return authSpec.AccountResponse{
		Email:            &entity.Email,
		FirstName:        &entity.FirstName,
		Id:               &entity.ID,
		LastName:         &entity.LastName,
		TelegramId:       &entity.TelegramId,
		TelegramUserName: &entity.TelegramUserName,
		Username:         &entity.Username,
	}
}

func (presenter *AccountPresenter) ToEntity(requestEntity *authSpec.CreateAccountRequest) domain.Account {
	return domain.Account{
		FirstName: *requestEntity.FirstName,
		LastName:  *requestEntity.LastName,
		Email:     *requestEntity.Email,
	}
}
