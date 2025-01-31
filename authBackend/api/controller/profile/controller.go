package profile

import (
	"authServer/api/presenter/profile"
	daoImpl "authServer/internal/dao"
	useCases "authServer/internal/usecases"
	"authServer/pkg/utilities"
	"github.com/gin-gonic/gin"
)

type ProfileController struct {
	ProfileUseCase *useCases.ProfileUseCase
	AccountUseCase *useCases.AccountUseCase
	Presenter      *profile.Presenter
}

func Create(dao daoImpl.AuthDao) *ProfileController {
	return &ProfileController{
		ProfileUseCase: &useCases.ProfileUseCase{Repo: dao.ProfileRepo},
		AccountUseCase: &useCases.AccountUseCase{Repo: dao.AccountRepo},
		Presenter:      &profile.Presenter{},
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
