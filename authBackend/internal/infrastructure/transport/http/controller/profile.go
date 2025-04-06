package controller

import (
	daoImpl "authServer/internal/app/ports/out/dao"
	"authServer/internal/app/services"
	useCases "authServer/internal/app/usecases"
	"authServer/internal/infrastructure/transport/http/presenter"
	"authServer/pkg/utilities"
	"github.com/gin-gonic/gin"
)

type ProfileController struct {
	profileUseCase useCases.GetProfileUCase
	getAccByTgId   *services.GetAccountUCaseImpl
	Presenter      *presenter.ProfilePresenter
}

func CreateProfileController(dao *daoImpl.AuthDao) *ProfileController {
	return &ProfileController{
		profileUseCase: &services.GetProfileUseCaseImpl{Repo: dao.ProfileRepo},
		getAccByTgId:   &services.GetAccountUCaseImpl{Repo: dao.AccountRepo},
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
