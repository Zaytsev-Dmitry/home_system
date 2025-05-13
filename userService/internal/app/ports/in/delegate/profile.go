package delegate

import (
	"userService/internal/app/domain"
	"userService/internal/app/ports/out/dao"
	"userService/internal/app/services"
	useCases "userService/internal/app/usecases"
)

type ProfileDelegate struct {
	createProfileUCase useCases.CreateProfileUCase
	getProfileUCase    useCases.GetProfileUCase
}

func (d *ProfileDelegate) Create(acc domain.Account) error {
	return d.createProfileUCase.Create(acc)
}

func (d *ProfileDelegate) GetByTGId(accId int64) (*domain.Profile, error) {
	return d.getProfileUCase.GetByTGId(accId)
}

func CreateProfileDelegate(dao *dao.AuthDao) *ProfileDelegate {
	return &ProfileDelegate{
		createProfileUCase: &services.CreateProfileUseCaseImpl{Repo: dao.ProfileRepository},
		getProfileUCase:    &services.GetProfileUseCaseImpl{Repo: dao.ProfileRepository},
	}
}
