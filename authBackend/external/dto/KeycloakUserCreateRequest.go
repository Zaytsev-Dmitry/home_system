package externalDto

import apiDto "authServer/api/docs"

type KeycloakUserCreateRequest struct {
	Enabled     bool          `json:"enabled"`
	Username    string        `json:"username"`
	Email       string        `json:"email"`
	FirstName   string        `json:"firstName"`
	LastName    string        `json:"lastName"`
	Credentials []Credentials `json:"credentials"`
}

type Credentials struct {
	Type      string `json:"type"`
	Value     string `json:"value"`
	Temporary bool   `json:"temporary"`
}

func NewKeycloakUserCreateRequest(request apiDto.CreateAccountRequest) *KeycloakUserCreateRequest {
	credentials := make([]Credentials, 1)
	credentials[0] = Credentials{
		Type: "password", Value: *request.Password, Temporary: false,
	}
	return &KeycloakUserCreateRequest{
		Enabled:     true,
		Username:    *request.Login,
		Email:       *request.Email,
		FirstName:   *request.FirstName,
		LastName:    *request.LastName,
		Credentials: credentials,
	}
}
