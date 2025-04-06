package delegate

import useCases "authBackend/internal/app/usecases"

type AccountDelegate struct {
	regAccountUCase useCases.RegisterAccountUseCase
	getAccountUCase useCases.GetAccountUCase
}
