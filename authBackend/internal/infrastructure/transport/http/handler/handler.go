package handler

import (
	daoImpl "authServer/internal/app/ports/out/dao"
	"authServer/internal/app/ports/out/keycloak"
	"authServer/internal/infrastructure/transport/http/controller"
	"authServer/pkg/config_loader"
	"github.com/gin-gonic/gin"
)

type AuthServerApi struct {
	accController     *controller.AccountController
	profileController *controller.ProfileController
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

func NewAuthServerApi(config *config_loader.AppConfig, dao *daoImpl.AuthDao) *AuthServerApi {
	keycloak := keycloak.New(config)
	return &AuthServerApi{
		accController:     controller.CreateAccountController(keycloak, dao),
		profileController: controller.CreateProfileController(dao),
	}
}
