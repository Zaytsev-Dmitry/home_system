package marshalling

import (
	"authBackend/pkg/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleMarshalling[T any](c *gin.Context, req *T) error {
	if err := c.ShouldBindJSON(req); err != nil {
		errors.SetResponseError(c, errors.MarshallError)
		return err
	}
	return nil
}

func HandleResponse[T any, R any](c *gin.Context, logic func() (T, error), present func(T, *gin.Context) R) {
	if result, err := logic(); err != nil {
		errors.HandleError(c, err)
	} else {
		c.JSON(http.StatusOK, present(result, c))
	}
}

func ExecuteAndRespondStatus(status int, c *gin.Context, logic func() error) {
	if err := logic(); err != nil {
		errors.HandleError(c, err)
	} else {
		c.Status(status)
	}
}
