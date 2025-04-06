package services

import (
	"authServer/internal/app/domain"
	"authServer/internal/dao/repository/intefraces"
)

type CreateUseCaseImpl struct {
	Repo intefraces.ProfileRepository
}

func (p *CreateUseCaseImpl) Create(acc domain.Account) error {
	return p.Repo.CreateProfile(acc)
}
