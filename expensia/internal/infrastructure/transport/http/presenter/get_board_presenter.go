package presenter

import (
	"expensia/api/openapi"
	"expensia/internal/app/domain"
	"github.com/gin-gonic/gin"
	"time"
)

type GetBoardPresenter struct {
}

func (c *GetBoardPresenter) Present(domains []*domain.Board, context *gin.Context) openapi.ListBoardBackendResponse {
	var boardsResp []openapi.BoardResponse
	for _, d := range domains {
		boardsResp = append(boardsResp, *c.toBoardResponse(d))
	}
	return openapi.ListBoardBackendResponse{
		Meta:    c.toMetaData(context),
		Payload: &boardsResp,
	}
}

func (c *GetBoardPresenter) toBoardResponse(domain *domain.Board) *openapi.BoardResponse {
	return &openapi.BoardResponse{
		Currency: &domain.Currency,
		Name:     &domain.Name,
		Owner:    &domain.OwnerId,
	}
}

func (c *GetBoardPresenter) toMetaData(context *gin.Context) *openapi.MetaData {
	nowString := time.Now().String()
	return &openapi.MetaData{
		Path:      &context.Request.URL.Path,
		Timestamp: &nowString,
	}
}
