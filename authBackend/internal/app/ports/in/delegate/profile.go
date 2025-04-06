package delegate

import (
	"authBackend/internal/app/domain"
	"authBackend/internal/app/ports/out/dao"
	"authBackend/internal/app/services"
	useCases "authBackend/internal/app/usecases"
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
