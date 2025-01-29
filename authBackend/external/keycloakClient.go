package external

import (
	externalDto "authServer/external/dto"
	"authServer/pkg/utilities"
	"fmt"
	authSpec "github.com/Zaytsev-Dmitry/home_system_open_api/authServerBackend"
	"net/http"
	"net/url"
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

func (client KeycloakClient) RegisterAccount(request authSpec.CreateAccountRequest) error {
	_, err := utilities.PostWithBearerAuthorization(
		client.getToken().AccessToken,
		externalDto.NewKeycloakUserCreateRequest(request),
		client.KeycloakHost+"/admin/realms/"+client.KeycloakRealm+"/users",
	)
	if err != nil {
		return err
	}
	return nil
}

func (client KeycloakClient) getToken() externalDto.KeycloakTokenResponse {
	data := url.Values{}
	data.Set("client_id", client.ClientId)
	data.Set("client_secret", client.ClientSecret)
	data.Set("grant_type", client.ServerGrantType)
	response := utilities.UrlencodedRequest(http.MethodPost, client.KeycloakUrl+client.TokenUrl, data)
	fmt.Println(response)
	var dto externalDto.KeycloakTokenResponse
	utilities.ParseResponseToStruct(response, &dto)
	return dto
}
