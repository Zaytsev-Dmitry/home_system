package controller

import (
	"authBackend/internal/app/ports/in/delegate"
	"authBackend/internal/app/ports/out/dao"
	"github.com/gin-gonic/gin"
)

type ProfileController struct {
	delegate *delegate.ProfileDelegate
}

func (c *ProfileController) GetProfileByTgId(context *gin.Context, telegramId int64) {
	c.delegate.GetByTGId(telegramId)
}

func CreateProfileController(dao *dao.AuthDao) *ProfileController {
	return &ProfileController{
		delegate: delegate.CreateProfileDelegate(dao),
	}
}
