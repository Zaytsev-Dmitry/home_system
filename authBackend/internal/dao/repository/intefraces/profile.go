package intefraces

import authServerDomain "authServer/internal/domain"

type ProfileRepository interface {
	CreateProfile(account authServerDomain.Account, tgUsername string) error
	GetProfileByAccountId(accId int64) authServerDomain.Profile
	CloseConnection()
}
