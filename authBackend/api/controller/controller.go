package noteControllers

import (
	presenter "authServer/api/presenter"
	"authServer/external"
	daoImpl "authServer/internal/dao"
	useCases "authServer/internal/usecases"
	"authServer/pkg/utilities"
	authSpec "github.com/Zaytsev-Dmitry/home_system_open_api/authServerBackend"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AccountController struct {
	RegisterUseCase *useCases.RegisterAccountUseCase
	presenter       *presenter.Presenter
}

// TODO надо отловить все кейсы keycloak
func (controller *AccountController) RegisterAccount(context *gin.Context) {
	var requestEntity authSpec.CreateAccountRequest
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

func Create(keycloakClient external.KeycloakClient, dao daoImpl.AuthDao) *AccountController {
	return &AccountController{
		RegisterUseCase: &useCases.RegisterAccountUseCase{
			Keycloak:       &keycloakClient,
			Repo:           dao.AccountRepo,
			ProfileUsecase: useCases.ProfileUseCase{Repo: dao.ProfileRepo},
		},
		presenter: &presenter.Presenter{},
	}
}
