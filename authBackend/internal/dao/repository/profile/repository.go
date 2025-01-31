package profile

import authServerDomain "authServer/internal/domain"

type ProfileRepository interface {
	CreateProfile(account authServerDomain.Account, tgUsername string) authServerDomain.Profile
	GetProfileByAccountId(accId int64) authServerDomain.Profile
	CloseConnection()
}
