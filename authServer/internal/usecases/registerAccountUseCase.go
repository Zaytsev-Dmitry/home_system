package usecases

import (
	apiDto "authServer/api/docs"
	keycloak "authServer/external"
	authDaoInterface "authServer/internal/dao/interface"
	domain "authServer/internal/domain"
)

type RegisterAccountUseCase struct {
	Keycloak *keycloak.KeycloakClient
	Dao      authDaoInterface.AuthDao
}

func (register *RegisterAccountUseCase) Register(request apiDto.CreateAccountRequest) (result domain.Account, err error) {
	return register.Dao.Save(
		domain.Account{
			FirstName: request.FirstName,
			LastName:  request.LastName,
			Email:     request.Email,
			Login:     *request.Login,
		},
	), nil
	//err = register.Keycloak.RegisterAccount(request)
	//if err != nil {
	//	return domain.Account{}, err
	//} else {
	//	return register.Dao.Save(
	//		domain.Account{
	//			FirstName: *request.FirstName,
	//			LastName:  *request.LastName,
	//			Email:     *request.Email,
	//			Login:     *request.Login,
	//		},
	//	), nil
	//}
}
