package services

import (
	"userService/internal/app/domain"
	"userService/internal/app/ports/out/dao/repository/profile"
)

type GetProfileUseCaseImpl struct {
	Repo profile.ProfileRepository
}

func (p *GetProfileUseCaseImpl) GetByTGId(accId int64) (*domain.Profile, error) {
	return p.Repo.GetProfileByAccountId(accId)
}
