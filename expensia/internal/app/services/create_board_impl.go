package services

import (
	"expensia/api/rest"
	"expensia/internal/app/domain"
	"expensia/internal/app/ports/out/dao/repository"
)

type CreateBoardUCaseImpl struct {
	Repo *repository.BoardRepository
}

func (c CreateBoardUCaseImpl) Create(params rest.CreateBoardParams) (*domain.Board, error) {
	return nil, nil
}
