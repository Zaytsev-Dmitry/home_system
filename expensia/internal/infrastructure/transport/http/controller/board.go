package controller

import (
	"expensia/internal/app/ports/in/delegate"
	"expensia/internal/app/ports/out/dao"
	"github.com/gin-gonic/gin"
)

type BoardController struct {
	delegate *delegate.BoardDelegate
}

func (bc BoardController) GetAllBoards(c *gin.Context) {
	//call delegate
}

func Create(dao *dao.ExpensiaDao) *BoardController {
	return &BoardController{
		delegate: &delegate.BoardDelegate{},
	}
}
