package controller

import (
	"expensia/api/rest"
	"expensia/internal/app/domain"
	"expensia/internal/app/ports/in/delegate"
	"expensia/internal/app/ports/out/dao"
	"expensia/internal/app/ports/out/dao/repository"
	"expensia/internal/app/prepare"
	"expensia/internal/infrastructure/transport/http/presenter"
	apikitHandler "github.com/Zaytsev-Dmitry/apikit/handlers"
	"github.com/gin-gonic/gin"
)

type BoardController struct {
	delegate          *delegate.BoardDelegate
	presenter         presenter.CreateBoardPresenter
	getBoardPresenter presenter.GetBoardPresenter
}

func (bc BoardController) GetAllBoards(context *gin.Context, tgUserId int64) {
	apikitHandler.HandleResponse(context, func() ([]domain.Board, error) {
		return bc.delegate.All(tgUserId)
	}, bc.getBoardPresenter.Present)
}

func (bc BoardController) CreateBoard(context *gin.Context, params rest.CreateBoardParams) {
	apikitHandler.HandleResponse(context, func() (*domain.Board, error) {
		return bc.delegate.CreateAndReturnBoard(
			repository.CreateBoardUCaseIn{
				Name:     params.Name,
				Currency: string(params.Currency),
			},
		)
	}, bc.presenter.Present)

}

func Create(dao *dao.ExpensiaDao, registry *prepare.PrepareRegistry) *BoardController {
	return &BoardController{
		delegate:  delegate.CreateBoardDelegate(dao, registry),
		presenter: presenter.CreateBoardPresenter{},
	}
}
