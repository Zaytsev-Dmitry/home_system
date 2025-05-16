package identity

import (
	"github.com/Nerzal/gocloak/v13"
	"userService/internal/app/domain"
)

type UserIdentityLinkRepository interface {
	Save(user *gocloak.User, tgId *int64) (*domain.UserIdentityLink, error)
	//CreateAccountAndProfile(entity keycloak.KeycloakEntity, username string, telegramId int64) (domain.Account, error)
	//GetIdByTgId(tgId int64) int64
	//GetByTgId(tgId int64) (domain.Account, error)
}
