package handler

import (
	"expensia/api/rest"
	"expensia/internal/app/ports/out/dao"
	"expensia/internal/infrastructure/transport/http/controller"
	"github.com/gin-gonic/gin"
)

type ExpensiaApi struct {
	boardController *controller.BoardController
}

func (api ExpensiaApi) GetAllBoards(c *gin.Context) {
	api.boardController.GetAllBoards(c)
}

func (api ExpensiaApi) CreateBoard(c *gin.Context, params rest.CreateBoardParams) {

}

func NewExpensiaApi(dao *dao.ExpensiaDao) *ExpensiaApi {
	return &ExpensiaApi{
		boardController: controller.Create(dao),
	}
}
