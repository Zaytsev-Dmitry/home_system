package external

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	authSpec "github.com/Zaytsev-Dmitry/home_system_open_api/authServerBackend"
	"net/http"
	"strconv"
	"telegramCLient/external/dto"
	"telegramCLient/pkg/utilities"
	"telegramCLient/util"
)

var (
	UNKNOWN   = errors.New("Unknown error. Can`t take account from keycloak")
	NOT_FOUND = errors.New("404")
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
func (serverClient *AuthServerClient) GetAccountByTgId(tgId int64) (dto.AccountDTO, error) {
	response, err := get(serverClient.AuthServerUrl + "/account/" + strconv.FormatInt(tgId, 10))
	if err != nil {
		return dto.AccountDTO{}, UNKNOWN
	}
	if response.StatusCode != http.StatusOK {
		if response.StatusCode == http.StatusNotFound {
			return dto.AccountDTO{}, NOT_FOUND
		}
		return dto.AccountDTO{}, UNKNOWN
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
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		err = utilities.Fail(utilities.Error{Err: err, Msg: "Cant create GET request"})
	}
	req.Header.Add("Accept", "application/json")
	do, err := client.Do(req)
	if err != nil {
		err = utilities.Fail(utilities.Error{Err: err, Msg: "Cant DO request"})
	}
	return do, err
}
