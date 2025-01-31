package noteHandlers

import (
	accountController "authServer/api/controller"
	authConfig "authServer/configs"
	"authServer/external"
	daoImpl "authServer/internal/dao"
	"github.com/gin-gonic/gin"
)

type AuthServerApi struct {
	controller *accountController.AccountController
}

func (api *AuthServerApi) RegisterAccount(context *gin.Context) {
	api.controller.RegisterAccount(context)
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
	return &AuthServerApi{controller: accountController.Create(keycloak, dao)}
}
