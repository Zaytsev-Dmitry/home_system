package keycloak

import (
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

func (client KeycloakClient) RegisterAccount(request authSpec.CreateAccountRequest) (error, KeycloakEntity) {
	var result KeycloakEntity
	serviceAccountToken, err := client.getToken()
	if err != nil {
		return errors.Join(Internal, errors.New("Wrap error: "+err.Error())), KeycloakEntity{}
	}
	resp, err := utilities.PostWithBearerAuthorization(
		serviceAccountToken.AccessToken,
		newKeycloakUserCreateRequest(request),
		client.KeycloakHost+"/admin/realms/"+client.KeycloakRealm+"/users",
	)
	if err != nil {
		return errors.Join(Internal, errors.New("Wrap error: "+err.Error())), KeycloakEntity{}
	}

	err = client.catchHttpStatus(resp)
	if err != nil {
		return err, KeycloakEntity{}
	}

	utilities.ParseResponseToStruct(resp, &result)
	defer resp.Body.Close()
	return err, result
}
func (client KeycloakClient) GetUser(mail string) (KeycloakEntity, error) {
	result := make([]KeycloakEntity, 0)

	var resultErr error
	serviceAccountToken, err := client.getToken()
	if err != nil {
		resultErr = errors.Join(Internal, errors.New("Wrap error: "+err.Error()))
	}
	resp, err := utilities.GetWithBearerAuthorization(
		serviceAccountToken.AccessToken,
		client.KeycloakHost+"/admin/realms/"+client.KeycloakRealm+"/users?email="+mail,
	)
	if err != nil {
		resultErr = errors.Join(Internal, err)
	}

	utilities.ParseResponseToStruct(resp, &result)
	return result[0], resultErr
}
