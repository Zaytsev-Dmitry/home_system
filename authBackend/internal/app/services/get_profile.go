package services

import (
	"authBackend/internal/app/domain"
	"authBackend/internal/app/ports/out/dao/repository/profile"
)

type GetProfileUseCaseImpl struct {
	Repo profile.ProfileRepository
}

func (p *GetProfileUseCaseImpl) GetByTGId(accId int64) domain.Profile {
	return p.Repo.GetProfileByAccountId(accId)
}
