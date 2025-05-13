package controller

import (
	"github.com/gin-gonic/gin"
	generatedApi "userService/api/http"
	"userService/internal/app/domain"
	"userService/internal/app/ports/in/delegate"
	daoImpl "userService/internal/app/ports/out/dao"
	"userService/internal/app/ports/out/keycloak"
	"userService/internal/infrastructure/transport/http/presenter"
	"userService/pkg/marshalling"
)

type AccountController struct {
	delegate  *delegate.AccountDelegate
	presenter presenter.AccountPresenter
}

func (cntr *AccountController) GetAccountByTgId(telegramId int64) {
	cntr.delegate.GetAccountIdByTgId(telegramId)
}

func (cntr *AccountController) RegisterAccount(context *gin.Context) {
	var req generatedApi.CreateAccountRequest
	if err := marshalling.HandleMarshalling(context, &req); err != nil {
		return
	}
	marshalling.HandleResponse(context, func() (domain.Account, error) {
		return cntr.delegate.Register(req)
	}, cntr.presenter.PresentToSingleResp)
}

func CreateAccountController(keycloakClient *keycloak.KeycloakClient, dao *daoImpl.AuthDao) *AccountController {
	return &AccountController{
		delegate:  delegate.CreateAccountDelegate(dao, *keycloakClient),
		presenter: presenter.AccountPresenter{},
	}
}
