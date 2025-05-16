package presenter

import (
	"expensia/api/rest"
	"expensia/internal/app/domain"
	"github.com/gin-gonic/gin"
	"time"
)

type GetBoardPresenter struct {
}

func (c *GetBoardPresenter) Present(domains []domain.Board, context *gin.Context) rest.ListBoardBackendResponse {
	var boardsResp []rest.BoardResponse
	for _, d := range domains {
		boardsResp = append(boardsResp, *c.toBoardResponse(d))
	}
	return rest.ListBoardBackendResponse{
		Meta:    c.toMetaData(context),
		Payload: &boardsResp,
	}
}

func (c *GetBoardPresenter) toBoardResponse(domain domain.Board) *rest.BoardResponse {
	return &rest.BoardResponse{
		Currency: &domain.Currency,
		Name:     &domain.Name,
		Owner:    &domain.OwnerId,
	}
}

func (c *GetBoardPresenter) toMetaData(context *gin.Context) *rest.MetaData {
	nowString := time.Now().String()
	return &rest.MetaData{
		Path:      &context.Request.URL.Path,
		Timestamp: &nowString,
	}
}
