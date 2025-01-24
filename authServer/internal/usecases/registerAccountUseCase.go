package usecases

import (
	apiDto "authServer/api/docs"
	keycloak "authServer/external"
	authDaoInterface "authServer/internal/dao/interface"
	domain "authServer/internal/domain"
	"net/http"
)

type RegisterAccountUseCase struct {
	Keycloak *keycloak.KeycloakClient
	Dao      authDaoInterface.AuthDao
}

func (register *RegisterAccountUseCase) Register(request apiDto.CreateAccountRequest) (result domain.Account, err error) {
	status := register.Keycloak.RegisterAccount(request)
	switch status {
	case http.StatusOK:
		return register.Dao.Save(
			domain.Account{
				FirstName: request.FirstName,
				LastName:  request.LastName,
				Email:     request.Email,
				Login:     *request.Login,
			},
		), nil
	case http.StatusConflict, http.StatusInternalServerError:
		return domain.Account{}, err
	default:
		return domain.Account{}, err
	}
}
