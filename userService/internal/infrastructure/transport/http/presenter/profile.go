package presenter

import (
	"github.com/gin-gonic/gin"
	"time"
	"userService/api/http"
	"userService/internal/app/domain"
)

type ProfilePresenter struct {
}

func (p *ProfilePresenter) PresentToSingleResp(domain *domain.Profile, context *gin.Context) http.SingleProfileBackendResponse {
	return http.SingleProfileBackendResponse{
		Meta:    p.toMetaData(context),
		Payload: p.toProfileResponse(domain),
	}
}

func (p *ProfilePresenter) toProfileResponse(note *domain.Profile) *http.ProfileResponse {
	return &http.ProfileResponse{}
}

func (p *ProfilePresenter) toMetaData(context *gin.Context) *http.MetaData {
	nowString := time.Now().String()
	return &http.MetaData{
		Path:      &context.Request.URL.Path,
		Timestamp: &nowString,
	}
}
