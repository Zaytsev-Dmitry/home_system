package keycloak

import (
	"authServer/pkg/utilities"
	authSpec "github.com/Zaytsev-Dmitry/home_system_open_api/authServerBackend"
	"net/http"
	"net/url"
)

type KeycloakTokenResponse struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        int64  `json:"expires_in"`
	RefreshExpiresIn int64  `json:"refresh_expires_in"`
	TokenType        string `json:"token_type"`
	NotBeforePolicy  int    `json:"not-before-policy"`
	Scope            string `json:"scope"`
}

type KeycloakUserCreateRequest struct {
	Enabled     bool          `json:"enabled"`
	Username    string        `json:"username"`
	Email       string        `json:"email"`
	FirstName   string        `json:"firstName"`
	LastName    string        `json:"lastName"`
	Credentials []Credentials `json:"credentials"`
}

type KeycloakEntity struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Enabled   bool   `json:"enabled"`
}

type Credentials struct {
	Type      string `json:"type"`
	Value     string `json:"value"`
	Temporary bool   `json:"temporary"`
}

type KeycloakResponseError struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
	ErrorMessage     string `json:"errorMessage"`
}

func newKeycloakUserCreateRequest(request authSpec.CreateAccountRequest) *KeycloakUserCreateRequest {
	credentials := make([]Credentials, 1)
	credentials[0] = Credentials{
		Type: "password", Value: *request.Password, Temporary: false,
	}
	return &KeycloakUserCreateRequest{
		Enabled:     true,
		Username:    *request.TelegramUsername,
		Email:       *request.Email,
		FirstName:   *request.FirstName,
		LastName:    *request.LastName,
		Credentials: credentials,
	}
}

func (client KeycloakClient) getToken() (KeycloakTokenResponse, error) {
	var result KeycloakTokenResponse
	data := url.Values{}
	data.Set("client_id", client.ClientId)
	data.Set("client_secret", client.ClientSecret)
	data.Set("grant_type", client.ServerGrantType)
	response, err := utilities.UrlencodedRequest(http.MethodPost, client.KeycloakUrl+client.TokenUrl, data)
	if err != nil {
		return result, err
	}
	utilities.ParseResponseToStruct(response, &result)
	return result, nil
}
