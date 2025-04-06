package account

import (
	"authBackend/internal/app/domain"
	"authBackend/internal/app/ports/out/keycloak"
)

type AccountRepository interface {
	CreateAccountAndProfile(entity keycloak.KeycloakEntity, username string, telegramId int64) (domain.Account, error)
	GetIdByTgId(tgId int64) int64
	GetByTgId(tgId int64) (domain.Account, error)
}
