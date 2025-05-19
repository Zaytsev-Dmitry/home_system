package services

import (
	"expensia/internal/app/domain"
	"expensia/internal/app/ports/out/dao/repository"
	dbkit "github.com/Zaytsev-Dmitry/dbkit/custom_error"
)

type GetBoardUCaseImpl struct {
	Repo repository.BoardRepository
}

func (c GetBoardUCaseImpl) GetAllBoards(ownerId int64) ([]*domain.Board, error) {
	res, err := c.Repo.GetAllByTgUserId(ownerId)
	if err != nil {
		return nil, dbkit.UpdateErrorText(err.Action, err.WrapErr)
	}
	return res, nil
}
