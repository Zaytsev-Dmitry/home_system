package controller

import (
	"github.com/gin-gonic/gin"
	"userService/internal/app/domain"
	"userService/internal/app/ports/in/delegate"
	"userService/internal/app/ports/out/dao"
	"userService/internal/infrastructure/transport/http/presenter"
	"userService/pkg/marshalling"
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
