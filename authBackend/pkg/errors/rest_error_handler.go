package errors

import (
	openapi "authBackend/api/http"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

var (
	RowNotFound     = errors.New("запись не найдена")
	ValidationError = errors.New("ошибка валидации")
	MarshallError   = errors.New("ошибка маршалинга")
)

func HandleError(c *gin.Context, err error) {
	log.Printf(err.Error())
	status, msg := getErrorMsgAndStatus(err)
	responseError := getErrorDto(msg, status, c)
	c.JSON(status, responseError)
}

func getErrorMsgAndStatus(err error) (int, string) {
	switch {
	case errors.Is(err, RowNotFound):
		return http.StatusNotFound, err.Error()
	case errors.Is(err, MarshallError), errors.Is(err, ValidationError):
		return http.StatusBadRequest, err.Error()
	default:
		return http.StatusInternalServerError, "Oops... что-то пошло не так"
	}
}

func getErrorDto(err string, errorCode int, context *gin.Context) openapi.BackendErrorResponse {
	nowString := time.Now().String()
	return openapi.BackendErrorResponse{
		Description: &err,
		ErrorCode:   &errorCode,
		Meta: &openapi.MetaData{
			Path:      &context.Request.URL.Path,
			Timestamp: &nowString,
		},
	}
}

func SetResponseError(context *gin.Context, err error) {
	var status int
	msg := "Oops...что то пошло не так"

	switch {
	case errors.Is(err, RowNotFound):
		status = http.StatusNotFound
		msg = err.Error()
	case errors.Is(err, MarshallError) || errors.Is(err, ValidationError):
		status = http.StatusBadRequest
		msg = err.Error()
	default:
		status = http.StatusInternalServerError
	}

	var responseError = getErrorDto(msg, status, context)
	context.Writer.Header().Set("Content-Type", "application/json")
	context.Status(status)
	json.NewEncoder(context.Writer).Encode(responseError)
}
