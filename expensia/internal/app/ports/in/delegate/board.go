package delegate

import (
	"expensia/internal/app/domain"
	"expensia/internal/app/ports/out/dao"
	"expensia/internal/app/ports/out/dao/repository"
	"expensia/internal/app/prepare"
	"expensia/internal/app/prepare/boardprepare"
	"expensia/internal/app/services"
	"expensia/internal/app/usecases"
)

type BoardDelegate struct {
	create   usecases.CreateBoardUCase
	get      usecases.GetBoardUCase
	registry *prepare.PrepareRegistry
}

func CreateBoardDelegate(dao *dao.ExpensiaDao, registry *prepare.PrepareRegistry) *BoardDelegate {
	boardprepare.RegisterBoardPreparer(registry, dao.ParticipantRepo)
	boardprepare.RegisterGetBoardsPreparer(registry, dao.ParticipantRepo)

	return &BoardDelegate{
		create:   services.CreateBoardUCaseImpl{Repo: dao.BoardRepo},
		get:      services.GetBoardUCaseImpl{Repo: dao.BoardRepo},
		registry: registry,
	}
}

func (d BoardDelegate) CreateAndReturnBoard(req repository.CreateBoardUCaseIn) (*domain.Board, error) {
	result, err := d.registry.Prepare("create_board", req)
	if err != nil {
		return nil, err
	}
	return d.create.CreateAndReturnBoard(result.(repository.CreateBoardUCaseIn))
}

func (d BoardDelegate) All(tgUserId int64) ([]domain.Board, error) {
	result, err := d.registry.Prepare("get_boards", tgUserId)
	if err != nil {
		return nil, err
	}
	return d.get.All(result.(int64))
}
