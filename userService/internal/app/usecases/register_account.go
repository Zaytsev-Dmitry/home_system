package usecases

import (
	generatedApi "userService/api/http"
	"userService/internal/app/domain"
)

type RegisterAccountUseCase interface {
	Register(request generatedApi.CreateAccountRequest) (domain.Account, error)
}
