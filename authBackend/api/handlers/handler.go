package handlers

import (
	accountController "authServer/api/controller"
	authConfig "authServer/configs"
	"authServer/external/keycloak"
	daoImpl "authServer/internal/dao"
	"github.com/gin-gonic/gin"
)

type AuthServerApi struct {
	accController     *accountController.AccountController
	profileController *accountController.ProfileController
}

func (api *AuthServerApi) GetAccountByTgId(c *gin.Context, telegramId int64) {
	api.accController.GetAccountByTgId(c, telegramId)
}

func (api *AuthServerApi) GetProfileByTgId(c *gin.Context, telegramId int64) {
	api.profileController.GetProfileByTgId(c, telegramId)
}

func (api *AuthServerApi) RegisterAccount(context *gin.Context) {
	api.accController.RegisterAccount(context)
}

func NewAuthServerApi(config *authConfig.AppConfig, dao daoImpl.AuthDao) *AuthServerApi {
	keycloak := keycloak.KeycloakClient{
		KeycloakUrl:     config.Keycloak.KeycloakUrl,
		TokenUrl:        config.Keycloak.TokenUrl,
		KeycloakHost:    config.Keycloak.KeycloakHost,
		KeycloakRealm:   config.Keycloak.KeycloakRealm,
		ClientId:        config.Keycloak.ClientId,
		ClientSecret:    config.Keycloak.ClientSecret,
		ServerGrantType: config.Keycloak.ServerGrantType,
	}
	return &AuthServerApi{
		accController:     accountController.CreateAccountController(keycloak, dao),
		profileController: accountController.CreateProfileController(dao),
	}
}
