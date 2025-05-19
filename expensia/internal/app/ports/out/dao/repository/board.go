package repository

import (
	"expensia/internal/app/domain"
)

type BoardRepository interface {
	SaveAndFlush(req CreateBoardUCaseIn) (*domain.Board, error)
	GetAllByTgUserId(ownerId int64) ([]domain.Board, error)
	GetById(boardId int64) (*domain.Board, error)
}
