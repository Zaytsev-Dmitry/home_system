package account

import authServerDomain "authServer/internal/domain"

type AccountRepository interface {
	Save(entity authServerDomain.Account) authServerDomain.Account
	GetIdByTgId(tgId int64) int64
	CloseConnection()
}
