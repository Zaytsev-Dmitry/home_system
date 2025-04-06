package delegate

import (
	generatedApi "authBackend/api/http"
	"authBackend/internal/app/domain"
	"authBackend/internal/app/ports/out/dao"
	"authBackend/internal/app/ports/out/keycloak"
	"authBackend/internal/app/services"
	useCases "authBackend/internal/app/usecases"
)

type AccountDelegate struct {
	regAccountUCase useCases.RegisterAccountUseCase
	getAccountUCase useCases.GetAccountUCase
}

func (cd *AccountDelegate) Register(request generatedApi.CreateAccountRequest) (domain.Account, error) {
	return cd.regAccountUCase.Register(request)
}

func (cd *AccountDelegate) Get(telegramId int64) (domain.Account, error) {
	return cd.getAccountUCase.Get(telegramId)
}

func (cd *AccountDelegate) GetAccountIdByTgId(tgId int64) (accId int64) {
	return cd.getAccountUCase.GetAccountIdByTgId(tgId)
}

func CreateAccountDelegate(dao *dao.AuthDao, client keycloak.KeycloakClient) *AccountDelegate {
	return &AccountDelegate{
		regAccountUCase: &services.RegisterAccountUseCaseImpl{
			Keycloak: &client,
			Repo:     dao.AccountRepository},
		getAccountUCase: &services.GetAccountUCaseImpl{Repo: dao.AccountRepository},
	}
}
