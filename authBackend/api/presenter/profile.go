package presenter

import (
	generatedApi "authServer/api/http"
	"authServer/internal/app/domain"
)

type ProfilePresenter struct {
}

func (presenter *ProfilePresenter) ToProfileResponse(entity domain.Profile) generatedApi.ProfileResponse {
	accId64 := int64(entity.AccountId)
	id64 := int64(entity.ID)
	return generatedApi.ProfileResponse{
		AccountId: &accId64,
		Id:        &id64,
		Role:      &entity.Role,
		Username:  &entity.Username,
	}
}
