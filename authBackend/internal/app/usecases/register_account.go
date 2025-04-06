package usecases

import (
	generatedApi "authBackend/api/http"
	"authBackend/internal/app/domain"
)

type RegisterAccountUseCase interface {
	Register(request generatedApi.CreateAccountRequest) (domain.Account, error)
}
