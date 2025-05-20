package repository

import (
	"expensia/internal/app/domain"
	"expensia/internal/app/usecases"
	"github.com/Zaytsev-Dmitry/dbkit/custom_error"
)

type BoardRepository interface {
	SaveAndFlush(req usecases.CreateBoardInput) (*domain.Board, error)
	GetAllByTgUserId(ownerId int64) ([]*domain.Board, *custom_error.CustomError)
	GetById(boardId int64) (*domain.Board, error)
}
