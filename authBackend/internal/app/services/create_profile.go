package services

import (
	"authBackend/internal/app/domain"
	"authBackend/internal/app/ports/out/dao/repository/profile"
)

type CreateProfileUseCaseImpl struct {
	Repo profile.ProfileRepository
}

func (p *CreateProfileUseCaseImpl) Create(acc domain.Account) error {
	return p.Repo.CreateProfile(acc)
}
