package controller

import (
	"authServer/api/presenter"
	"authServer/internal/app/ports/out/keycloak"
	"authServer/internal/app/services"
	useCases "authServer/internal/app/usecases"
	daoImpl "authServer/internal/dao"
	"github.com/gin-gonic/gin"
)

type AccountController struct {
	registerAcc  useCases.RegisterAccountUseCase
	getAccByTgId *services.GetAccountUCaseImpl
	presenter    *presenter.AccountPresenter
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
		registerAcc: &services.RegisterAccountUseCaseImpl{
			Keycloak: keycloakClient,
			Repo:     dao.AccountRepo,
		},
		getAccByTgId: &services.GetAccountUCaseImpl{
			Repo: dao.AccountRepo,
		},
		presenter: &presenter.AccountPresenter{},
	}
}
