package handler

import (
	"expensia/api/openapi"
	"expensia/internal/app/ports/out/dao"
	"expensia/internal/app/prepare"
	"expensia/internal/infrastructure/transport/http/controller"
	"github.com/gin-gonic/gin"
)

type ExpensiaApi struct {
	boardController *controller.BoardController
}

func (api ExpensiaApi) GetAllBoards(c *gin.Context, params openapi.GetAllBoardsParams) {
	api.boardController.GetAllBoards(c, params)
}

func (api ExpensiaApi) CreateBoard(c *gin.Context, params openapi.CreateBoardParams) {
	api.boardController.CreateBoard(c, params)
}

func (api ExpensiaApi) AddParticipantToBoard(c *gin.Context, params openapi.AddParticipantToBoardParams) {
	api.boardController.AddParticipantToBoard(c, params)
}

func NewExpensiaApi(dao *dao.ExpensiaDao, registry *prepare.PrepareRegistry) *ExpensiaApi {
	return &ExpensiaApi{
		boardController: controller.Create(dao, registry),
	}
}
