package noteHandlers

import (
	accountController "authServer/api/controller/account"
	profileController "authServer/api/controller/profile"
	authConfig "authServer/configs"
	"authServer/external"
	daoImpl "authServer/internal/dao"
	"github.com/gin-gonic/gin"
)

type AuthServerApi struct {
	accController     *accountController.AccountController
	profileController *profileController.ProfileController
}

func (api *AuthServerApi) GetProfileByTgId(c *gin.Context, telegramId int) {
	api.profileController.GetProfileByTgId(c, int64(telegramId))
}

func (api *AuthServerApi) RegisterAccount(context *gin.Context) {
	api.accController.RegisterAccount(context)
}

func NewAuthServerApi(config *authConfig.AppConfig, dao daoImpl.AuthDao) *AuthServerApi {
	keycloak := external.KeycloakClient{
		KeycloakUrl:     config.Keycloak.KeycloakUrl,
		TokenUrl:        config.Keycloak.TokenUrl,
		KeycloakHost:    config.Keycloak.KeycloakHost,
		KeycloakRealm:   config.Keycloak.KeycloakRealm,
		ClientId:        config.Keycloak.ClientId,
		ClientSecret:    config.Keycloak.ClientSecret,
		ServerGrantType: config.Keycloak.ServerGrantType,
	}
	return &AuthServerApi{
		accController:     accountController.Create(keycloak, dao),
		profileController: profileController.Create(dao),
	}
}
