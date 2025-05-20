package delegate

import (
	"expensia/internal/app/domain"
	"expensia/internal/app/ports/out/dao"
	"expensia/internal/app/ports/out/dao/repository"
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
		CreateBoard:           createBoardFunc(dao.ParticipantRepo, dao.BoardRepo),
		AddParticipantToBoard: addParticipantsFunc(dao.ParticipantRepo, dao.BoardRepo, dao.BoardParticipantRepo),
		GetBoardsByTgId:       getBoardsFunc(dao.ParticipantRepo, dao.BoardRepo),
	}
}

func createBoardFunc(partRepo repository.ParticipantRepository, boardRepo repository.BoardRepository) func(usecases.CreateBoardInput) (*domain.Board, error) {
	return func(in usecases.CreateBoardInput) (*domain.Board, error) {
		prepared, err := prepare.PrepareCreateBoard(in, partRepo)
		if err != nil {
			return nil, err
		}
		return services.CreateBoardUCaseImpl{Repo: boardRepo}.CreateAndReturnBoard(prepared)
	}
}

func addParticipantsFunc(partRepo repository.ParticipantRepository, boardRepo repository.BoardRepository, boardPartRepo repository.BoardParticipantRepository) func(usecases.AddParticipantsInput) error {
	return func(in usecases.AddParticipantsInput) error {
		prepared, err := prepare.PrepareAddParticipants(in, partRepo, boardRepo)
		if err != nil {
			return err
		}
		return services.AddParticipantUCaseImpl{Repo: boardPartRepo}.AddParticipantsToBoard(prepared)
	}
}

func getBoardsFunc(partRepo repository.ParticipantRepository, boardRepo repository.BoardRepository) func(int64) ([]*domain.Board, error) {
	return func(tgUserId int64) ([]*domain.Board, error) {
		id, err := prepare.PrepareGetBoards(tgUserId, partRepo)
		if err != nil {
			return nil, err
		}
		return services.GetBoardUCaseImpl{Repo: boardRepo}.GetAllBoards(id)
	}
}
