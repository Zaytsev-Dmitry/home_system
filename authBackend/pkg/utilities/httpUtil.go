package utilities

import (
	"bytes"
	"encoding/json"
	"errors"
	authSpec "github.com/Zaytsev-Dmitry/home_system_open_api/authServerBackend"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var client = &http.Client{}

var (
	MarshallError      = errors.New("marshalling error")
	HttpCreateReqError = errors.New("new request creating error")
	HttpDoRequestError = errors.New("client do error")
	ParseRequest       = errors.New("parse request URI error")
)

func ParseResponseToStruct(respBody *http.Response, response any) any {
	defer respBody.Body.Close()
	decoder := json.NewDecoder(respBody.Body)
	decoder.Decode(response)
	return response
}

func PostWithBearerAuthorization(token string, body any, url string) (*http.Response, error) {
	var err error
	marshal, err := json.Marshal(body)
	if err != nil {
		err = MarshallError
	}
	req, err := http.NewRequest("POST", url, bytes.NewReader(marshal))
	if err != nil {
		err = HttpCreateReqError
	}
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Accept", "application/json")
	do, err := client.Do(req)
	if err != nil {
		err = HttpDoRequestError
	}
	return do, err
}

func UrlencodedRequest(httpMethod string, urlStr string, data url.Values) (*http.Response, error) {
	uri, err := url.ParseRequestURI(urlStr)
	if err != nil {
		err = ParseRequest
	}
	r, err := http.NewRequest(httpMethod, uri.String(), strings.NewReader(data.Encode()))
	if err != nil {
		err = HttpCreateReqError
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(r)
	if err != nil {
		err = HttpDoRequestError
	}
	return resp, err
}

func getErrorDto(err string, status int, context *gin.Context) authSpec.ErrorResponse {
	nowString := time.Now().String()
	return authSpec.ErrorResponse{
		Timestamp: &nowString,
		Status:    &status,
		Error:     &err,
		Path:      &context.Request.URL.Path,
	}
}

func CatchMarshallErr(err error, context *gin.Context) {
	if err != nil {
		var responseError = getErrorDto(err.Error(), http.StatusBadRequest, context)
		context.Status(http.StatusBadRequest)
		json.NewEncoder(context.Writer).Encode(responseError)
	}
}

func SetResponse(v any, context *gin.Context) {
	json.NewEncoder(context.Writer).Encode(v)
}

func SetResponseError(context *gin.Context, status int) {
	var responseError = getErrorDto("Oops...что то пошло не так", status, context)
	context.Writer.Header().Set("Content-Type", "application/json")
	context.Status(status)
	json.NewEncoder(context.Writer).Encode(responseError)
}
