package intefraces

import (
	authServerDomain "authServer/internal/app/domain"
	"authServer/internal/app/ports/out/keycloak"
)

type AccountRepository interface {
	CreateAccountAndProfile(entity keycloak.KeycloakEntity, username string, telegramId int64) (authServerDomain.Account, error)
	GetIdByTgId(tgId int64) int64
	GetByTgId(tgId int64) (authServerDomain.Account, error)
	CloseConnection()
}
