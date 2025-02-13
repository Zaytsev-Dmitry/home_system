package usecases

import (
	"authServer/internal/dao/repository"
	"authServer/pkg/utilities"
	"errors"
	"net/http"
)

func IfExistErrLogAndReturn500Http(respErr error, status int) int {
	if respErr != nil {
		utilities.GetLogger().Error(respErr.Error())
		if errors.Is(respErr, repository.NoRows) {
			status = http.StatusNotFound
		} else {
			status = http.StatusInternalServerError
		}
	}
	return status
}
