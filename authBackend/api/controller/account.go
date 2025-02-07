package controller

import (
	"authServer/api/presenter"
	"authServer/external/keycloak"
	daoImpl "authServer/internal/dao"
	useCases "authServer/internal/usecases"
	"authServer/pkg/utilities"
	authSpec "github.com/Zaytsev-Dmitry/home_system_open_api/authServerBackend"
	"github.com/gin-gonic/gin"
)

type AccountController struct {
	RegisterUseCase *useCases.AccountUseCase
	presenter       *presenter.AccountPresenter
}

func (controller *AccountController) RegisterAccount(context *gin.Context) {
	var requestEntity authSpec.CreateAccountRequest
	utilities.CatchMarshallErr(context.BindJSON(&requestEntity), context)
	entity, err, status := controller.RegisterUseCase.Register(requestEntity)
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
		RegisterUseCase: &useCases.AccountUseCase{
			Keycloak: &keycloakClient,
			Repo:     dao.AccountRepo,
		},
		presenter: &presenter.AccountPresenter{},
	}
}
