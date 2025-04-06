package services

import (
	"authServer/internal/app/domain"
	"authServer/internal/dao/repository/intefraces"
)

type GetProfileUseCaseImpl struct {
	Repo intefraces.ProfileRepository
}

func (p *GetProfileUseCaseImpl) GetByTGId(accId int64) domain.Profile {
	return p.Repo.GetProfileByAccountId(accId)
}
