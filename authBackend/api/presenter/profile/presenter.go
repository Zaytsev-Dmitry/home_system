package profile

import (
	domain "authServer/internal/domain"
	authSpec "github.com/Zaytsev-Dmitry/home_system_open_api/authServerBackend"
)

type Presenter struct {
}

func (presenter *Presenter) ToProfileResponse(entity domain.Profile) authSpec.ProfileResponse {
	accId64 := int64(entity.AccountId)
	id64 := int64(entity.ID)
	return authSpec.ProfileResponse{
		AccountId:        &accId64,
		Id:               &id64,
		Role:             &entity.Role,
		TelegramUsername: &entity.TelegramUsername,
	}
}
