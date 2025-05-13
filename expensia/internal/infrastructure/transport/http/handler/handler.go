package handler

import (
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

func NewExpensiaApi(dao *dao.ExpensiaDao) *ExpensiaApi {
	return &ExpensiaApi{
		boardController: controller.Create(dao),
	}
}
