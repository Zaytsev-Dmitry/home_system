package account

import authServerDomain "authServer/internal/domain"

type AccountRepository interface {
	Save(entity authServerDomain.Account) authServerDomain.Account
	CloseConnection()
}
