package account

import (
	"userService/internal/app/domain"
	"userService/internal/app/ports/out/keycloak"
)

type AccountRepository interface {
	CreateAccountAndProfile(entity keycloak.KeycloakEntity, username string, telegramId int64) (domain.Account, error)
	GetIdByTgId(tgId int64) int64
	GetByTgId(tgId int64) (domain.Account, error)
}
