package usecases

import "authServer/internal/app/domain"

type GetProfileUCase interface {
	GetByTGId(accId int64) domain.Profile
}
