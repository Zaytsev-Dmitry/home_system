package presenter

import (
	"github.com/gin-gonic/gin"
	"time"
	"userService/api/http"
	"userService/internal/app/domain"
)

type AccountPresenter struct {
}

func (p *AccountPresenter) PresentToSingleResp(domain *domain.UserIdentityLink, context *gin.Context) http.SingleUserIdentityBackendResponse {
	return http.SingleUserIdentityBackendResponse{
		Meta:    p.toMetaData(context),
		Payload: p.toUserIdentityResponse(domain),
	}
}

func (p *AccountPresenter) toUserIdentityResponse(obj *domain.UserIdentityLink) *http.UserIdentityResponse {
	return &http.UserIdentityResponse{
		Email:          &obj.Email,
		TelegramUserId: &obj.TelegramUserID,
	}
}

func (p *AccountPresenter) toMetaData(context *gin.Context) *http.MetaData {
	nowString := time.Now().String()
	return &http.MetaData{
		Path:      &context.Request.URL.Path,
		Timestamp: &nowString,
	}
}
