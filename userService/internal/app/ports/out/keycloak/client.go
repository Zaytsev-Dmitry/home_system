package keycloak

import (
	"errors"
	openapi "userService/api/http"
	"userService/pkg/config_loader"
	"userService/pkg/utilities"
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

func New(config *config_loader.AppConfig) *KeycloakClient {
	return &KeycloakClient{
		KeycloakUrl:     config.Keycloak.KeycloakUrl,
		TokenUrl:        config.Keycloak.TokenUrl,
		KeycloakHost:    config.Keycloak.KeycloakHost,
		KeycloakRealm:   config.Keycloak.KeycloakRealm,
		ClientId:        config.Keycloak.ClientId,
		ClientSecret:    config.Keycloak.ClientSecret,
		ServerGrantType: config.Keycloak.ServerGrantType,
	}
}

func (client KeycloakClient) RegisterAccount(request openapi.CreateAccountRequest) (error, KeycloakEntity) {
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
	defer resp.Body.Close()
	return err, result
}
func (client KeycloakClient) GetUser(mail string) (KeycloakEntity, error) {
	result := make([]KeycloakEntity, 0)

	var resultErr error
	serviceAccountToken, err := client.getToken()
	if err != nil {
		resultErr = errors.Join(Internal, errors.New("Wrap error: "+err.Error()))
		return KeycloakEntity{}, resultErr
	}

	resp, err := utilities.GetWithBearerAuthorization(
		serviceAccountToken.AccessToken,
		client.KeycloakHost+"/admin/realms/"+client.KeycloakRealm+"/users?email="+mail,
	)

	if err != nil {
		resultErr = errors.Join(Internal, errors.New("Wrap error: "+err.Error()))
	}

	utilities.ParseResponseToStruct(resp, &result)
	return result[0], resultErr
}
