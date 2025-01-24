package authDaoInterface

import authServerDomain "authServer/internal/domain"

type AuthDao interface {
	Save(entity authServerDomain.Account) authServerDomain.Account
}
