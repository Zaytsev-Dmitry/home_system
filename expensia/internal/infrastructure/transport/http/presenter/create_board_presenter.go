package presenter

import (
	"expensia/api/rest"
	"expensia/internal/app/domain"
	"github.com/gin-gonic/gin"
	"time"
)

type CreateBoardPresenter struct {
}

func (c *CreateBoardPresenter) Present(domain *domain.Board, context *gin.Context) rest.CreateBoardBackendResponse {
	return rest.CreateBoardBackendResponse{
		Meta:    c.toMetaData(context),
		Payload: c.toBoardResponse(domain),
	}
}

func (c *CreateBoardPresenter) toBoardResponse(domain *domain.Board) *rest.BoardResponse {
	return &rest.BoardResponse{
		Currency: &domain.Currency,
		Name:     &domain.Name,
		Owner:    &domain.OwnerId,
	}
}

func (c *CreateBoardPresenter) toMetaData(context *gin.Context) *rest.MetaData {
	nowString := time.Now().String()
	return &rest.MetaData{
		Path:      &context.Request.URL.Path,
		Timestamp: &nowString,
	}
}
