package repository

import (
	"expensia/internal/app/domain"
	"github.com/Zaytsev-Dmitry/dbkit/custom_error"
)

type BoardRepository interface {
	SaveAndFlush(req CreateBoardUCaseIn) (*domain.Board, error)
	GetAllByTgUserId(ownerId int64) ([]*domain.Board, *custom_error.CustomError)
	GetById(boardId int64) (*domain.Board, error)
}
