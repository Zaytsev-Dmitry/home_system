package intefraces

import (
	"authServer/external/keycloak"
	authServerDomain "authServer/internal/domain"
)

type AccountRepository interface {
	CreateAccountAndProfile(entity keycloak.KeycloakEntity, telegramId int) (authServerDomain.Account, error)
	GetIdByTgId(tgId int64) int64
	GetByTgId(tgId int64) (authServerDomain.Account, error)
	CloseConnection()
}
