package controller

import (
	"authServer/api/presenter"
	daoImpl "authServer/internal/dao"
	useCases "authServer/internal/usecases"
	"authServer/pkg/utilities"
	"github.com/gin-gonic/gin"
)

type ProfileController struct {
	profileUseCase *useCases.ProfileUseCase
	getAccByTgId   *useCases.GetAccByTelegramId
	Presenter      *presenter.ProfilePresenter
}

func CreateProfileController(dao *daoImpl.AuthDao) *ProfileController {
	return &ProfileController{
		profileUseCase: &useCases.ProfileUseCase{Repo: dao.ProfileRepo},
		getAccByTgId:   &useCases.GetAccByTelegramId{Repo: dao.AccountRepo},
		Presenter:      &presenter.ProfilePresenter{},
	}
}

func (controller *ProfileController) GetProfileByTgId(context *gin.Context, telegramId int64) {
	accId := controller.getAccByTgId.GetAccountIdByTgId(telegramId)
	profileEntity := controller.profileUseCase.GetByTGId(accId)
	response := controller.Presenter.ToProfileResponse(profileEntity)
	utilities.SetResponse(
		response,
		context,
	)
}
