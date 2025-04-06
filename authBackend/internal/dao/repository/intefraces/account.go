package intefraces

import (
	"authServer/external/keycloak"
	authServerDomain "authServer/internal/app/domain"
)

type AccountRepository interface {
	CreateAccountAndProfile(entity keycloak.KeycloakEntity, username string, telegramId int64) (authServerDomain.Account, error)
	GetIdByTgId(tgId int64) int64
	GetByTgId(tgId int64) (authServerDomain.Account, error)
	CloseConnection()
}
