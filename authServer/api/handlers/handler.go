package noteHandlers

import (
	accountController "authServer/api/controller"
	"github.com/gin-gonic/gin"
)

type AuthServerApi struct {
	controller *accountController.AccountController
}

func (api *AuthServerApi) RegisterAccount(context *gin.Context) {
	api.controller.RegisterAccount(context)
}

func NewAuthServerApi() *AuthServerApi {
	return &AuthServerApi{controller: accountController.Create()}
}
