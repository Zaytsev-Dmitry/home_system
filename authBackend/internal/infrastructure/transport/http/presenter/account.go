package presenter

import (
	"authBackend/api/http"
	"authBackend/internal/app/domain"
	"github.com/gin-gonic/gin"
	"time"
)

type AccountPresenter struct {
}

func (p *AccountPresenter) PresentToSingleResp(domain *domain.Account, context *gin.Context) http.SingleAccountBackendResponse {
	return http.SingleAccountBackendResponse{
		Meta:    p.toMetaData(context),
		Payload: p.toAccountResponse(domain),
	}
}

func (p *AccountPresenter) toAccountResponse(note *domain.Account) *http.AccountResponse {
	return &http.AccountResponse{}
}

func (p *AccountPresenter) toMetaData(context *gin.Context) *http.MetaData {
	nowString := time.Now().String()
	return &http.MetaData{
		Path:      &context.Request.URL.Path,
		Timestamp: &nowString,
	}
}
