package controller

import (
	"expensia/api/rest"
	"expensia/internal/app/ports/in/delegate"
	"expensia/internal/app/ports/out/dao"
	"github.com/gin-gonic/gin"
)

type BoardController struct {
	delegate *delegate.BoardDelegate
}

func (bc BoardController) GetAllBoards(c *gin.Context) {
	bc.delegate.All()
}

func (bc BoardController) CreateBoard(c *gin.Context, params rest.CreateBoardParams) {
	bc.delegate.Create(params)
}

func Create(dao *dao.ExpensiaDao) *BoardController {
	return &BoardController{
		delegate: delegate.CreateBoardDelegate(dao),
	}
}
