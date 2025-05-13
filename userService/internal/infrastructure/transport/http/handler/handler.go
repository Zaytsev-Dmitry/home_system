package handler

import (
	"github.com/gin-gonic/gin"
	daoImpl "userService/internal/app/ports/out/dao"
	"userService/internal/app/ports/out/keycloak"
	"userService/internal/infrastructure/transport/http/controller"
	"userService/pkg/config_loader"
)

type AuthServerApi struct {
	accController     *controller.AccountController
	profileController *controller.ProfileController
}

func (api *AuthServerApi) GetAccountByTgId(c *gin.Context, telegramId int64) {
	api.accController.GetAccountByTgId(telegramId)
}

func (api *AuthServerApi) GetProfileByTgId(c *gin.Context, telegramId int64) {
	api.profileController.GetProfileByTgId(c, telegramId)
}

func (api *AuthServerApi) RegisterAccount(context *gin.Context) {
	api.accController.RegisterAccount(context)
}

func NewAuthServerApi(config *config_loader.AppConfig, dao *daoImpl.AuthDao) *AuthServerApi {
	return &AuthServerApi{
		accController:     controller.CreateAccountController(keycloak.New(config), dao),
		profileController: controller.CreateProfileController(dao),
	}
}
