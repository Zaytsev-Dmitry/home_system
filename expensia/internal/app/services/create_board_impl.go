package services

import (
	"expensia/internal/app/domain"
	"expensia/internal/app/ports/out/dao/repository"
)

type CreateBoardUCaseImpl struct {
	Repo repository.BoardRepository
}

func (c CreateBoardUCaseImpl) CreateAndReturnBoard(req repository.CreateBoardUCaseIn) (*domain.Board, error) {
	return c.Repo.SaveAndFlush(req)
}
