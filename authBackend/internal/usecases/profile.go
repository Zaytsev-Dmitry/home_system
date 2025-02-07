package usecases

import (
	"authServer/internal/dao/repository/intefraces"
	domain "authServer/internal/domain"
)

type ProfileUseCase struct {
	Repo intefraces.ProfileRepository
}

func (p *ProfileUseCase) Create(acc domain.Account, tgUsername string) error {
	return p.Repo.CreateProfile(acc, tgUsername)
}

func (p *ProfileUseCase) GetByTGId(accId int64) domain.Profile {
	return p.Repo.GetProfileByAccountId(accId)
}
