package utilities

import (
	apiDTO "authServer/api/docs"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var client = &http.Client{}

func UrlencodedRequest(httpMethod string, urlStr string, data url.Values) *http.Response {
	uri, _ := url.ParseRequestURI(urlStr)
	r, _ := http.NewRequest(httpMethod, uri.String(), strings.NewReader(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, _ := client.Do(r)
	return resp
}

func getErrorDto(err string, status int, context *gin.Context) apiDTO.ErrorResponse {
	nowString := time.Now().String()
	return apiDTO.ErrorResponse{
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

func SetResponseError(err error, context *gin.Context, status int) {
	var responseError = getErrorDto(err.Error(), status, context)
	context.Status(status)
	json.NewEncoder(context.Writer).Encode(responseError)
}
