package usecases

import (
	generatedApi "userService/api/http"
	"userService/internal/app/domain"
)

type RegisterUserUseCase interface {
	Register(request generatedApi.CreateAccountRequest) (*domain.UserIdentityLink, error)
}
