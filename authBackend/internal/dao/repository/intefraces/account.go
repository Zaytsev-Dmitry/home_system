package intefraces

import authServerDomain "authServer/internal/domain"

type AccountRepository interface {
	Save(entity authServerDomain.Account) (authServerDomain.Account, error)
	Register(entity authServerDomain.Account) (authServerDomain.Account, error)
	GetIdByTgId(tgId int64) int64
	GetByTgId(tgId int64) (authServerDomain.Account, error)
	CloseConnection()
}
