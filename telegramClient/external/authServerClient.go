package external

import (
	"bytes"
	"encoding/json"
	"net/http"
	"telegramCLient/external/dto"
	"telegramCLient/util"
)

var client = &http.Client{}

type AuthServerClient struct {
	AuthServerUrl string
}

func NewAuthServerClient(authServerUrl string) *AuthServerClient {
	return &AuthServerClient{authServerUrl}
}

// TODO отловаить ошибки
func (serverClient *AuthServerClient) RegisterUser(source dto.CreateAccountRequest) dto.AccountDTO {
	response, _ := post(source, serverClient.AuthServerUrl+"/account")
	var respDto dto.AccountResponse
	util.ParseResponseToStruct(response, &respDto)
	return dto.AccountDTO{
		FirstName: *respDto.Login,
		LastName:  *respDto.Login,
		Login:     *respDto.Login,
		Email:     *respDto.Login,
	}
}

func post(body any, url string) (*http.Response, error) {
	marshal, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", url, bytes.NewReader(marshal))
	req.Header.Add("Accept", "application/json")
	return client.Do(req)
}
