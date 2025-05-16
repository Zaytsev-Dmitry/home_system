package usecases

import (
	"expensia/internal/app/domain"
)

type GetBoardUCase interface {
	All() *[]domain.Board
}
