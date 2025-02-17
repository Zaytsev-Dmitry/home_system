package controller

import (
	"authServer/api/presenter"
	generatedApi "authServer/api/spec"
	"authServer/external/keycloak"
	daoImpl "authServer/internal/dao"
	"authServer/internal/domain"
	useCases "authServer/internal/usecases"
	"authServer/pkg/utilities"
	"github.com/gin-gonic/gin"
)

type AccountController struct {
	registerAcc  *useCases.RegisterAccount
	getAccByTgId *useCases.GetAccByTelegramId
	presenter    *presenter.AccountPresenter
}

func (controller *AccountController) RegisterAccount(context *gin.Context) {
	var requestEntity generatedApi.CreateAccountRequest
	utilities.CatchMarshallErr(context.BindJSON(&requestEntity), context)
	entity, err, status := controller.registerAcc.Register(requestEntity)
	controller.processAccountResult(context, err, status, entity)
}

func (controller *AccountController) GetAccountByTgId(context *gin.Context, telegramId int64) {
	entity, err, status := controller.getAccByTgId.Get(telegramId)
	controller.processAccountResult(context, err, status, entity)
}

func (controller *AccountController) processAccountResult(context *gin.Context, err error, status int, entity authServerDomain.Account) {
	if err != nil {
		utilities.SetResponseError(context, status)
	} else {
		utilities.SetResponseWithStatus(
			controller.presenter.ToAccountResponse(entity),
			context,
			status,
		)
	}
}

func CreateAccountController(keycloakClient keycloak.KeycloakClient, dao daoImpl.AuthDao) *AccountController {
	return &AccountController{
		registerAcc: &useCases.RegisterAccount{
			Keycloak: &keycloakClient,
			Repo:     dao.AccountRepo,
		},
		getAccByTgId: &useCases.GetAccByTelegramId{
			Repo: dao.AccountRepo,
		},
		presenter: &presenter.AccountPresenter{},
	}
}
