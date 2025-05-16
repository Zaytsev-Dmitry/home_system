package delegate

import (
	"go.uber.org/zap"
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

func (cd *UserDelegate) Register(request generatedApi.CreateAccountRequest, logger *zap.Logger) (*domain.UserIdentityLink, error) {
	logger.Info("Run register account delegate")
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
