package handlers

import (
	accountController "authServer/api/controller"
	"authServer/internal/app/ports/out/keycloak"
	daoImpl "authServer/internal/dao"
	"authServer/pkg/config_loader"
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

func NewAuthServerApi(config *config_loader.AppConfig, dao *daoImpl.AuthDao) *AuthServerApi {
	keycloak := keycloak.New(config)
	return &AuthServerApi{
		accController:     accountController.CreateAccountController(keycloak, dao),
		profileController: accountController.CreateProfileController(dao),
	}
}
