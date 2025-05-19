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
	create         usecases.CreateBoardUCase
	get            usecases.GetBoardUCase
	addParticipant usecases.AddParticipantUCase
	registry       *prepare.PrepareRegistry
}

func CreateBoardDelegate(dao *dao.ExpensiaDao, registry *prepare.PrepareRegistry) *BoardDelegate {
	boardprepare.RegisterCreateBoardPreparer(registry, dao.ParticipantRepo)
	boardprepare.RegisterGetBoardsPreparer(registry, dao.ParticipantRepo)
	boardprepare.RegisterAddParticipantPreparer(registry, dao.ParticipantRepo, dao.BoardRepo)

	return &BoardDelegate{
		create:         services.CreateBoardUCaseImpl{Repo: dao.BoardRepo},
		get:            services.GetBoardUCaseImpl{Repo: dao.BoardRepo},
		addParticipant: services.AddParticipantUCaseImpl{Repo: dao.BoardParticipantRepo},
		registry:       registry,
	}
}

func (d BoardDelegate) CreateAndReturnBoard(req repository.CreateBoardUCaseIn) (*domain.Board, error) {
	return prepare.WithPrepared(
		d.registry,
		"create_board",
		req,
		d.create.CreateAndReturnBoard,
	)
}

func (d BoardDelegate) AddParticipantsToBoard(req repository.AddParticipantsInput) error {
	return prepare.WithPreparedNoResult(
		d.registry,
		"add_participant_to_board",
		req,
		d.addParticipant.AddParticipantsToBoard,
	)
}

func (d BoardDelegate) GetAllBoards(tgUserId int64) ([]*domain.Board, error) {
	return prepare.WithPrepared(
		d.registry,
		"get_boards",
		tgUserId,
		d.get.GetAllBoards,
	)
}
