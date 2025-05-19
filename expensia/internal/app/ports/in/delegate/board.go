package delegate

import (
	"expensia/internal/app/domain"
	"expensia/internal/app/ports/out/dao"
	"expensia/internal/app/ports/out/dao/repository"
	"expensia/internal/app/prepare"
	"expensia/internal/app/prepare/boardprepare"
	"expensia/internal/app/services"
)

type BoardDelegate struct {
	CreateBoard           func(repository.CreateBoardUCaseIn) (*domain.Board, error)
	AddParticipantToBoard func(repository.AddParticipantsInput) error
	GetBoardsByTgId       func(int64) ([]*domain.Board, error)
}

func CreateBoardDelegate(dao *dao.ExpensiaDao, registry *prepare.PrepareRegistry) *BoardDelegate {
	boardprepare.RegisterCreateBoardPreparer(registry, dao.ParticipantRepo)
	boardprepare.RegisterGetBoardsPreparer(registry, dao.ParticipantRepo)
	boardprepare.RegisterAddParticipantPreparer(registry, dao.ParticipantRepo, dao.BoardRepo)

	return &BoardDelegate{
		CreateBoard: MakeDelegateWithResult(
			registry,
			"create_board",
			services.CreateBoardUCaseImpl{Repo: dao.BoardRepo}.CreateAndReturnBoard,
		),
		AddParticipantToBoard: MakeDelegateNoResult(
			registry,
			"add_participant_to_board",
			services.AddParticipantUCaseImpl{Repo: dao.BoardParticipantRepo}.AddParticipantsToBoard,
		),
		GetBoardsByTgId: MakeDelegateWithResult(
			registry,
			"get_boards",
			services.GetBoardUCaseImpl{Repo: dao.BoardRepo}.GetAllBoards,
		),
	}
}
