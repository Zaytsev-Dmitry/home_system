package controller

import (
	"authBackend/internal/app/ports/out/dao"
	"authBackend/internal/app/services"
	useCases "authBackend/internal/app/usecases"
	"github.com/gin-gonic/gin"
)

type ProfileController struct {
	profileUseCase useCases.GetProfileUCase
}

func (controller *ProfileController) GetProfileByTgId(context *gin.Context, telegramId int64) {
	controller.profileUseCase.GetByTGId(telegramId)
}

func CreateProfileController(dao *dao.AuthDao) *ProfileController {
	return &ProfileController{
		profileUseCase: &services.GetProfileUseCaseImpl{Repo: dao.ProfileRepository},
	}
}
