package usecases

import (
	"expensia/internal/app/domain"
	"expensia/internal/app/ports/out/dao/repository"
)

type CreateBoardUCase interface {
	CreateAndReturnBoard(req repository.CreateBoardUCaseIn) (*domain.Board, error)
}
