package handler

import (
	"github.com/gin-gonic/gin"
	daoImpl "userService/internal/app/ports/out/dao"
	"userService/internal/app/ports/out/keycloak"
	"userService/internal/infrastructure/transport/http/controller"
	"userService/pkg/config_loader"
)

type AuthServerApi struct {
	identityController *controller.IdentityUserController
}

func (api *AuthServerApi) GetAccountByTgId(c *gin.Context, telegramId int64) {

}

func (api *AuthServerApi) GetProfileByTgId(c *gin.Context, telegramId int64) {

}

func (api *AuthServerApi) RegisterAccount(context *gin.Context) {
	api.identityController.RegisterAccount(context)
}

func NewAuthServerApi(config *config_loader.AppConfig, dao *daoImpl.UserDao) *AuthServerApi {
	return &AuthServerApi{
		identityController: controller.CreateAccountController(keycloak.NewKeycloakClient(config), dao),
	}
}
