package services

import (
	"userService/internal/app/domain"
	"userService/internal/app/ports/out/dao/repository/profile"
)

type CreateProfileUseCaseImpl struct {
	Repo profile.ProfileRepository
}

func (p *CreateProfileUseCaseImpl) Create(acc domain.Account) error {
	return p.Repo.CreateProfile(acc)
}
