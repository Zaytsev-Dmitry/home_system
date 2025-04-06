package usecases

import "authBackend/internal/app/domain"

type GetProfileUCase interface {
	GetByTGId(accId int64) domain.Profile
}
