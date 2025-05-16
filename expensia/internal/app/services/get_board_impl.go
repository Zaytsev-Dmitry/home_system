package services

import (
	"expensia/internal/app/domain"
	"expensia/internal/app/ports/out/dao/repository"
)

type GetBoardUCaseImpl struct {
	Repo repository.BoardRepository
}

func (c GetBoardUCaseImpl) All(ownerId int64) ([]domain.Board, error) {
	return c.Repo.GetAllByTgUserId(ownerId)
}
