package usecases

import (
	"expensia/internal/app/domain"
)

type GetBoardUCase interface {
	GetAllBoards(ownerId int64) ([]domain.Board, error)
}
