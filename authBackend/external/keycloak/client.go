package keycloak

import (
	domain "authServer/internal/domain"
	"authServer/pkg/utilities"
	"errors"
	authSpec "github.com/Zaytsev-Dmitry/home_system_open_api/authServerBackend"
)

type KeycloakClient struct {
	KeycloakUrl     string
	KeycloakHost    string
	KeycloakRealm   string
	TokenUrl        string
	ClientId        string
	ClientSecret    string
	ServerGrantType string
}

func (client KeycloakClient) RegisterAccount(request authSpec.CreateAccountRequest) (error, domain.Account) {
	var result domain.Account
	serviceAccountToken, err := client.getToken()
	if err != nil {
		return errors.Join(Internal, errors.New("Wrap error: "+err.Error())), domain.Account{}
	}
	resp, err := utilities.PostWithBearerAuthorization(
		serviceAccountToken.AccessToken,
		newKeycloakUserCreateRequest(request),
		client.KeycloakHost+"/admin/realms/"+client.KeycloakRealm+"/users",
	)
	if err != nil {
		return errors.Join(Internal, errors.New("Wrap error: "+err.Error())), domain.Account{}
	}

	err = client.catchHttpStatus(resp)
	if err != nil {
		return err, domain.Account{}
	}

	utilities.ParseResponseToStruct(resp, &result)
	defer resp.Body.Close()
	return err, result
}
