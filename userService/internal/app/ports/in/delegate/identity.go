package delegate

import (
	generatedApi "userService/api/http"
	"userService/internal/app/domain"
	"userService/internal/app/ports/out/dao"
	"userService/internal/app/ports/out/keycloak"
	"userService/internal/app/services"
	useCases "userService/internal/app/usecases"
)

type UserDelegate struct {
	regUserUCase useCases.RegisterUserUseCase
}

func (cd *UserDelegate) Register(request generatedApi.CreateAccountRequest) (*domain.UserIdentityLink, error) {
	return cd.regUserUCase.Register(request)
}

func CreateAccountDelegate(dao *dao.UserDao, client keycloak.KeycloakClient) *UserDelegate {
	return &UserDelegate{
		regUserUCase: &services.RegisterAccountUseCaseImpl{
			Keycloak: &client,
			Repo:     dao.IdentityRepo,
		},
	}
}
