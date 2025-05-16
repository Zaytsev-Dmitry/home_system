package delegate

import (
	"expensia/api/rest"
	"expensia/internal/app/domain"
	"expensia/internal/app/ports/out/dao"
	"expensia/internal/app/services"
	"expensia/internal/app/usecases"
)

type BoardDelegate struct {
	create usecases.CreateBoardUCase
	get    usecases.GetBoardUCase
}

func CreateBoardDelegate(dao *dao.ExpensiaDao) *BoardDelegate {
	return &BoardDelegate{
		create: services.CreateBoardUCaseImpl{Repo: &dao.BoardRepo},
		get:    services.GetBoardUCaseImpl{Repo: &dao.BoardRepo},
	}
}

func (d BoardDelegate) Create(params rest.CreateBoardParams) (*domain.Board, error) {
	return nil, nil
}

func (d BoardDelegate) All() *[]domain.Board {
	return nil
}
