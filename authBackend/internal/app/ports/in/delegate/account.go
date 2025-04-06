package delegate

import useCases "authServer/internal/app/usecases"

type AccountDelegate struct {
	regAccountUCase useCases.RegisterAccountUseCase
	getAccountUCase useCases.GetAccountUCase
}
