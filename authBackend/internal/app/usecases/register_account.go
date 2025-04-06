package usecases

import (
	generatedApi "authServer/api/http"
	"authServer/internal/app/domain"
)

type RegisterAccountUseCase interface {
	Register(request generatedApi.CreateAccountRequest) (domain.Account, error)
}
