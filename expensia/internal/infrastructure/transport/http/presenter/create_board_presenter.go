package presenter

import (
	"expensia/api/openapi"
	"expensia/internal/app/domain"
	"github.com/gin-gonic/gin"
	"time"
)

type CreateBoardPresenter struct {
}

func (c *CreateBoardPresenter) Present(domain *domain.Board, context *gin.Context) openapi.SingleBoardBackendResponse {
	return openapi.SingleBoardBackendResponse{
		Meta:    c.toMetaData(context),
		Payload: c.toBoardResponse(domain),
	}
}

func (c *CreateBoardPresenter) toBoardResponse(domain *domain.Board) *openapi.BoardResponse {
	return &openapi.BoardResponse{
		Currency: &domain.Currency,
		Name:     &domain.Name,
		Owner:    &domain.OwnerId,
	}
}

func (c *CreateBoardPresenter) toMetaData(context *gin.Context) *openapi.MetaData {
	nowString := time.Now().String()
	return &openapi.MetaData{
		Path:      &context.Request.URL.Path,
		Timestamp: &nowString,
	}
}
