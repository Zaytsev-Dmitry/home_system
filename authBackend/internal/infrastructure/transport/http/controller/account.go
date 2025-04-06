package controller

import (
	daoImpl "authBackend/internal/app/ports/out/dao"
	"authBackend/internal/app/ports/out/keycloak"
	"authBackend/internal/app/services"
	useCases "authBackend/internal/app/usecases"
	"github.com/gin-gonic/gin"
)

type AccountController struct {
	register   useCases.RegisterAccountUseCase
	getAccount useCases.GetAccountUCase
}

func (controller *AccountController) RegisterAccount(context *gin.Context) {
	//var requestEntity generatedApi.CreateAccountRequest
	//utilities.CatchMarshallErr(context.BindJSON(&requestEntity), context)
	//entity, err := controller.registerAcc.Register(requestEntity)
	//controller.processAccountResult(context, err, entity)
}

func (controller *AccountController) GetAccountByTgId(context *gin.Context, telegramId int64) {
	//entity, err := controller.getAccByTgId.Get(telegramId)
	//controller.processAccountResult(context, err, entity)
}

func (controller *AccountController) processAccountResult(context *gin.Context, err error) {
	//if err != nil {
	//	utilities.SetResponseError(context, http.StatusInternalServerError)
	//} else {
	//	utilities.SetResponseWithStatus(
	//		controller.presenter.ToAccountResponse(entity),
	//		context,
	//		http.StatusOK,
	//	)
	//}
}

func CreateAccountController(keycloakClient *keycloak.KeycloakClient, dao *daoImpl.AuthDao) *AccountController {
	return &AccountController{
		register:   &services.RegisterAccountUseCaseImpl{Repo: dao.AccountRepository},
		getAccount: &services.GetAccountUCaseImpl{Repo: dao.AccountRepository},
	}
}
