package intefraces

import (
	authServerDomain "authServer/internal/app/domain"
)

type ProfileRepository interface {
	CreateProfile(account authServerDomain.Account) error
	GetProfileByAccountId(accId int64) authServerDomain.Profile
	CloseConnection()
}
