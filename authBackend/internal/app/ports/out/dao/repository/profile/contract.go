package profile

import (
	"authBackend/internal/app/domain"
)

type ProfileRepository interface {
	CreateProfile(account domain.Account) error
	GetProfileByAccountId(accId int64) domain.Profile
}
