package controller

import (
	"authBackend/internal/app/domain"
	"authBackend/internal/app/ports/in/delegate"
	"authBackend/internal/app/ports/out/dao"
	"authBackend/internal/infrastructure/transport/http/presenter"
	"authBackend/pkg/marshalling"
	"github.com/gin-gonic/gin"
)

type ProfileController struct {
	delegate  *delegate.ProfileDelegate
	presenter presenter.ProfilePresenter
}

func (c *ProfileController) GetProfileByTgId(context *gin.Context, telegramId int64) {
	marshalling.HandleResponse(context, func() (*domain.Profile, error) {
		return c.delegate.GetByTGId(telegramId)
	}, c.presenter.PresentToSingleResp)
}

func CreateProfileController(dao *dao.AuthDao) *ProfileController {
	return &ProfileController{
		delegate:  delegate.CreateProfileDelegate(dao),
		presenter: presenter.ProfilePresenter{},
	}
}
