package noteUtilities

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	noteApiDTO "noteBackendApp/api/docs"
	"time"
)

func getErrorDto(err string, status int, context *gin.Context) noteApiDTO.ErrorResponse {
	nowString := time.Now().String()
	return noteApiDTO.ErrorResponse{
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
