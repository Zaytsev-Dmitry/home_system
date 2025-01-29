package noteUtilities

import (
	"encoding/json"
	noteSpec "github.com/Zaytsev-Dmitry/home_system_open_api/noteServerBackend"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func getErrorDto(err string, status int, context *gin.Context) noteSpec.ErrorResponse {
	nowString := time.Now().String()
	return noteSpec.ErrorResponse{
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
