package noteUtilities

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type ErrorResponse struct {
	BusinessCode *int    `json:"businessCode,omitempty"`
	Error        *string `json:"error,omitempty"`
	Path         *string `json:"path,omitempty"`
	Status       *int    `json:"status,omitempty"`
	Timestamp    *string `json:"timestamp,omitempty"`
}

func getErrorDto(err string, status int, context *gin.Context) ErrorResponse {
	nowString := time.Now().String()
	return ErrorResponse{
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
