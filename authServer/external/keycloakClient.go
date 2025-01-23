package external

import (
	apiDto "authServer/api/docs"
	"authServer/pkg/utilities"
	"fmt"
	"net/http"
	"net/url"
)

type KeycloakClient struct {
	KeycloakUrl     string
	TokenUrl        string
	ClientId        string
	ClientSecret    string
	ServerGrantType string
}

func (client KeycloakClient) RegisterAccount(request apiDto.CreateAccountRequest) {
	client.getToken()
}

func (client KeycloakClient) getToken() {
	data := url.Values{}
	data.Set("client_id", client.ClientId)
	data.Set("client_secret", client.ClientSecret)
	data.Set("grant_type", client.ServerGrantType)
	fmt.Println(utilities.UrlencodedRequest(http.MethodPost, client.KeycloakUrl+client.TokenUrl, data))
}
