package usecases

import (
	"authServer/internal/dao/repository/profile"
	domain "authServer/internal/domain"
)

type ProfileUseCase struct {
	Repo profile.ProfileRepository
}

func (p *ProfileUseCase) Create(acc domain.Account, tgUsername string) {
	p.Repo.CreateProfile(acc, tgUsername)
}

func (p *ProfileUseCase) GetByTGId(accId int64) domain.Profile {
	return p.Repo.GetProfileByAccountId(accId)
}
