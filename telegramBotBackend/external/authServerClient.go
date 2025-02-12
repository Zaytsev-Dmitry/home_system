package external

import (
	"bytes"
	"encoding/json"
	"fmt"
	authSpec "github.com/Zaytsev-Dmitry/home_system_open_api/authServerBackend"
	"net/http"
	"strconv"
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
func (serverClient *AuthServerClient) RegisterUser(source authSpec.CreateAccountRequest) (dto.AccountDTO, error) {
	response, err := post(source, serverClient.AuthServerUrl+"/account")
	if err != nil {
		fmt.Println(err)
	}
	var respDto authSpec.AccountResponse
	util.ParseResponseToStruct(response, &respDto)
	return dto.AccountDTO{
		ID:         respDto.Id,
		FirstName:  respDto.FirstName,
		LastName:   respDto.LastName,
		Username:   respDto.Username,
		Email:      respDto.Email,
		TgUsername: respDto.TelegramUserName,
		TelegramId: respDto.TelegramId,
	}, err
}

// TODO отловаить ошибки
func (serverClient *AuthServerClient) GetAccountByTgId(tgId int64) dto.AccountDTO {
	response, err := get(serverClient.AuthServerUrl + "/account/" + strconv.FormatInt(tgId, 10))
	if err != nil {
		fmt.Println(err)
	}
	var respDto authSpec.AccountResponse
	util.ParseResponseToStruct(response, &respDto)
	return dto.AccountDTO{
		ID:         respDto.Id,
		FirstName:  respDto.FirstName,
		LastName:   respDto.LastName,
		Username:   respDto.Username,
		Email:      respDto.Email,
		TgUsername: respDto.TelegramUserName,
		TelegramId: respDto.TelegramId,
	}
}

// TODO отловаить ошибки
func (serverClient *AuthServerClient) GetProfileByTelegramId(tgId int) authSpec.ProfileResponse {
	response, err := get(serverClient.AuthServerUrl + "/profile/" + strconv.Itoa(tgId))
	if err != nil {
		fmt.Println(err)
	}
	var respDto authSpec.ProfileResponse
	util.ParseResponseToStruct(response, &respDto)
	return respDto
}

func post(body any, url string) (*http.Response, error) {
	marshal, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", url, bytes.NewReader(marshal))
	req.Header.Add("Accept", "application/json")
	return client.Do(req)
}

func get(url string) (*http.Response, error) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "application/json")
	return client.Do(req)
}
