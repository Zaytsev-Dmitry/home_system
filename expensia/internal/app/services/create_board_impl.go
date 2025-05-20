package services

import (
	"expensia/internal/app/domain"
	"expensia/internal/app/ports/out/dao/repository"
	"expensia/internal/app/usecases"
)

type CreateBoardUCaseImpl struct {
	Repo repository.BoardRepository
}

func (c CreateBoardUCaseImpl) CreateAndReturnBoard(req usecases.CreateBoardInput) (*domain.Board, error) {
	return c.Repo.SaveAndFlush(req)
}
