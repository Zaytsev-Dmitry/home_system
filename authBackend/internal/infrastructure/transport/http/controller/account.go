package controller

import (
	"authBackend/internal/app/ports/in/delegate"
	daoImpl "authBackend/internal/app/ports/out/dao"
	"authBackend/internal/app/ports/out/keycloak"
	"authBackend/internal/infrastructure/transport/http/presenter"
)

type AccountController struct {
	delegate  *delegate.AccountDelegate
	presenter presenter.AccountPresenter
}

func CreateAccountController(keycloakClient *keycloak.KeycloakClient, dao *daoImpl.AuthDao) *AccountController {
	return &AccountController{
		delegate:  delegate.CreateAccountDelegate(dao, *keycloakClient),
		presenter: presenter.AccountPresenter{},
	}
}
