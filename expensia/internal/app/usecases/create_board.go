package usecases

import (
	"expensia/api/rest"
	"expensia/internal/app/domain"
)

type CreateBoardUCase interface {
	Create(params rest.CreateBoardParams) (*domain.Board, error)
}
