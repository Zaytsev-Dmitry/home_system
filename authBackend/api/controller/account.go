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
	"net/http"
)

type AccountController struct {
	registerAcc  *useCases.RegisterAccount
	getAccByTgId *useCases.GetAccByTelegramId
	presenter    *presenter.AccountPresenter
}

func (controller *AccountController) RegisterAccount(context *gin.Context) {
	var requestEntity generatedApi.CreateAccountRequest
	utilities.CatchMarshallErr(context.BindJSON(&requestEntity), context)
	entity, err := controller.registerAcc.Register(requestEntity)
	controller.processAccountResult(context, err, entity)
}

func (controller *AccountController) GetAccountByTgId(context *gin.Context, telegramId int64) {
	entity, err := controller.getAccByTgId.Get(telegramId)
	controller.processAccountResult(context, err, entity)
}

func (controller *AccountController) processAccountResult(context *gin.Context, err error, entity authServerDomain.Account) {
	if err != nil {
		utilities.SetResponseError(context, http.StatusInternalServerError)
	} else {
		utilities.SetResponseWithStatus(
			controller.presenter.ToAccountResponse(entity),
			context,
			http.StatusOK,
		)
	}
}

func CreateAccountController(keycloakClient *keycloak.KeycloakClient, dao *daoImpl.AuthDao) *AccountController {
	return &AccountController{
		registerAcc: &useCases.RegisterAccount{
			Keycloak: keycloakClient,
			Repo:     dao.AccountRepo,
		},
		getAccByTgId: &useCases.GetAccByTelegramId{
			Repo: dao.AccountRepo,
		},
		presenter: &presenter.AccountPresenter{},
	}
}
