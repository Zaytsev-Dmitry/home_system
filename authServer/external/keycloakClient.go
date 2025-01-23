package external

import (
	apiDto "authServer/api/docs"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type KeycloakClient struct {
	KeycloakUrl     string
	TokenUrl        string
	ClientId        string
	ClientSecret    string
	ServerGrantType string
}

var client = &http.Client{}

func (client KeycloakClient) RegisterAccount(request apiDto.CreateAccountRequest) {
	client.getToken()
}

func (client KeycloakClient) getToken() {
	data := url.Values{}
	data.Set("client_id", client.ClientId)
	data.Set("client_secret", client.ClientSecret)
	data.Set("grant_type", client.ServerGrantType)
	fmt.Println(urlencodedRequest(http.MethodPost, client.KeycloakUrl+client.TokenUrl, data))
}

func urlencodedRequest(httpMethod string, urlStr string, data url.Values) *http.Response {
	uri, _ := url.ParseRequestURI(urlStr)
	r, _ := http.NewRequest(httpMethod, uri.String(), strings.NewReader(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, _ := client.Do(r)
	return resp
}
