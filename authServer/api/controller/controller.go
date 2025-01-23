package noteControllers

import (
	apiDTO "authServer/api/docs"
	presenter "authServer/api/presenter"
	"authServer/external"
	useCases "authServer/internal/usecases"
	"authServer/pkg/utilities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AccountController struct {
	RegisterUseCase *useCases.RegisterAccountUseCase
	presenter       *presenter.Presenter
}

func (controller *AccountController) RegisterAccount(context *gin.Context) {
	var requestEntity apiDTO.CreateAccountRequest
	utilities.CatchMarshallErr(context.BindJSON(&requestEntity), context)
	entity, err := controller.RegisterUseCase.Register(requestEntity)
	if err != nil {
		utilities.SetResponseError(err, context, http.StatusNotFound)
	}
	utilities.SetResponse(
		controller.presenter.ToAccountResponse(entity),
		context,
	)
}

func Create(keycloakClient external.KeycloakClient) *AccountController {
	return &AccountController{
		RegisterUseCase: &useCases.RegisterAccountUseCase{Keycloak: &keycloakClient},
		presenter:       &presenter.Presenter{},
	}
}
