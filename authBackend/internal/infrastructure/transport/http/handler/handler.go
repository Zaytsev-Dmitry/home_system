package handler

import (
	daoImpl "authBackend/internal/app/ports/out/dao"
	"authBackend/internal/app/ports/out/keycloak"
	"authBackend/internal/infrastructure/transport/http/controller"
	"authBackend/pkg/config_loader"
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
	return &AuthServerApi{
		accController:     controller.CreateAccountController(keycloak.New(config), dao),
		profileController: controller.CreateProfileController(dao),
	}
}
