package noteControllers

import (
	apiDTO "authServer/api/docs"
	presenter "authServer/api/presenter"
	"authServer/external"
	authDaoInterface "authServer/internal/dao/interface"
	useCases "authServer/internal/usecases"
	"authServer/pkg/utilities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AccountController struct {
	RegisterUseCase *useCases.RegisterAccountUseCase
	presenter       *presenter.Presenter
}

// TODO надо отловить все кейсы keycloak
func (controller *AccountController) RegisterAccount(context *gin.Context) {
	var requestEntity apiDTO.CreateAccountRequest
	utilities.CatchMarshallErr(context.BindJSON(&requestEntity), context)
	entity, err := controller.RegisterUseCase.Register(requestEntity)
	if err != nil {
		utilities.SetResponseError(err, context, http.StatusInternalServerError)
	}
	utilities.SetResponse(
		controller.presenter.ToAccountResponse(entity),
		context,
	)
}

func Create(keycloakClient external.KeycloakClient, dao authDaoInterface.AuthDao) *AccountController {
	return &AccountController{
		RegisterUseCase: &useCases.RegisterAccountUseCase{
			Keycloak: &keycloakClient,
			Dao:      dao,
		},
		presenter: &presenter.Presenter{},
	}
}
