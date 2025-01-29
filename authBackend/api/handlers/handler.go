package noteHandlers

import (
	accountController "authServer/api/controller"
	"authServer/external"
	authDaoInterface "authServer/internal/dao/interface"
	"github.com/gin-gonic/gin"
)

type AuthServerApi struct {
	controller *accountController.AccountController
}

func (api *AuthServerApi) RegisterAccount(context *gin.Context) {
	api.controller.RegisterAccount(context)
}

func NewAuthServerApi(keycloakClient external.KeycloakClient, dao authDaoInterface.AuthDao) *AuthServerApi {
	return &AuthServerApi{controller: accountController.Create(keycloakClient, dao)}
}
