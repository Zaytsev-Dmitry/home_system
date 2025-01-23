package authServerPresenter

import (
	apiDTO "authServer/api/docs"
	domain "authServer/internal/domain"
	"github.com/google/uuid"
)

type Presenter struct {
}

func (presenter *Presenter) ToAccountResponse(entity domain.AccountEntity) apiDTO.AccountResponse {
	return apiDTO.AccountResponse{
		Email:     &entity.Email,
		FirstName: &entity.FirstName,
		LastName:  &entity.LastName,
		Login:     &entity.Login,
	}
}

func (presenter *Presenter) ToEntity(requestEntity *apiDTO.CreateAccountRequest) domain.AccountEntity {
	return domain.AccountEntity{
		Id:        uuid.New().String(),
		FirstName: *requestEntity.FirstName,
		LastName:  *requestEntity.LastName,
		Login:     *requestEntity.Login,
		Email:     *requestEntity.Email,
	}
}
