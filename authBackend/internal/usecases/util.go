package usecases

import (
	"authServer/pkg/utilities"
	"net/http"
)

func IfExistErrLogAndReturn500Http(respErr error, status int) int {
	if respErr != nil {
		utilities.GetLogger().Error(respErr.Error())
		status = http.StatusInternalServerError
	}
	return status
}
