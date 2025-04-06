package controller

import (
	"authBackend/internal/app/ports/in/delegate"
	daoImpl "authBackend/internal/app/ports/out/dao"
	"authBackend/internal/app/ports/out/keycloak"
	"github.com/gin-gonic/gin"
)

type AccountController struct {
	delegate *delegate.AccountDelegate
}

func (cntr *AccountController) RegisterAccount(c *gin.Context) {
	//var req openapi.CreateAccountRequest
	//if err := marshalling.HandleMarshalling(c, &req); err != nil {
	//	return
	//}
	//marshalling.HandleResponse(c, func() (*domain.Account, error) {
	//	return cntr.delegate.Save(cntr.presenter.PresentToReq(req))
	//}, cntr.presenter.PresentToResp)

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
		delegate: delegate.CreateAccountDelegate(dao, *keycloakClient),
	}
}
