package usecases

import (
	"expensia/internal/app/domain"
)

type GetBoardUCase interface {
	All(ownerId int64) ([]domain.Board, error)
}
