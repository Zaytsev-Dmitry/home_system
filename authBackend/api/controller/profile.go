package controller

import (
	"authServer/api/presenter"
	daoImpl "authServer/internal/dao"
	useCases "authServer/internal/usecases"
	"authServer/pkg/utilities"
	"github.com/gin-gonic/gin"
)

type ProfileController struct {
	ProfileUseCase *useCases.ProfileUseCase
	AccountUseCase *useCases.AccountUseCase
	Presenter      *presenter.ProfilePresenter
}

func CreateProfileController(dao daoImpl.AuthDao) *ProfileController {
	return &ProfileController{
		ProfileUseCase: &useCases.ProfileUseCase{Repo: dao.ProfileRepo},
		AccountUseCase: &useCases.AccountUseCase{Repo: dao.AccountRepo},
		Presenter:      &presenter.ProfilePresenter{},
	}
}

func (controller *ProfileController) GetProfileByTgId(context *gin.Context, telegramId int64) {
	accId := controller.AccountUseCase.GetAccountIdByTgId(telegramId)
	profileEntity := controller.ProfileUseCase.GetByTGId(accId)
	response := controller.Presenter.ToProfileResponse(profileEntity)
	utilities.SetResponse(
		response,
		context,
	)
}
