package delegate

import (
	"expensia/internal/app/domain"
	"expensia/internal/app/ports/out/dao"
	"expensia/internal/app/prepare"
	"expensia/internal/app/services"
	"expensia/internal/app/usecases"
)

type BoardDelegate struct {
	CreateBoard           func(input usecases.CreateBoardInput) (*domain.Board, error)
	AddParticipantToBoard func(usecases.AddParticipantsInput) error
	GetBoardsByTgId       func(int64) ([]*domain.Board, error)
}

func CreateBoardDelegate(dao *dao.ExpensiaDao) *BoardDelegate {
	return &BoardDelegate{
		CreateBoard: func(in usecases.CreateBoardInput) (*domain.Board, error) {
			prepared, err := prepare.PrepareCreateBoard(in, dao.ParticipantRepo)
			if err != nil {
				return nil, err
			}
			return services.CreateBoardUCaseImpl{Repo: dao.BoardRepo}.CreateAndReturnBoard(prepared)
		},
		AddParticipantToBoard: func(in usecases.AddParticipantsInput) error {
			prepared, err := prepare.PrepareAddParticipants(in, dao.ParticipantRepo, dao.BoardRepo)
			if err != nil {
				return err
			}
			return services.AddParticipantUCaseImpl{Repo: dao.BoardParticipantRepo}.AddParticipantsToBoard(prepared)
		},
		GetBoardsByTgId: func(tgUserId int64) ([]*domain.Board, error) {
			id, err := prepare.PrepareGetBoards(tgUserId, dao.ParticipantRepo)
			if err != nil {
				return nil, err
			}
			return services.GetBoardUCaseImpl{Repo: dao.BoardRepo}.GetAllBoards(id)
		},
	}
}
