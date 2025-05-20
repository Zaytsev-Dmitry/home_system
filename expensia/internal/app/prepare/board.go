package prepare

import (
	"expensia/internal/app/ports/out/dao/repository"
	"expensia/internal/app/usecases"
)

func PrepareCreateBoard(in usecases.CreateBoardInput, repo repository.ParticipantRepository) (usecases.CreateBoardInput, error) {
	id, err := repo.GetIdByTgUserId(in.TgUserId)
	if err != nil {
		return usecases.CreateBoardInput{}, err
	}
	in.OwnerId = id
	return in, nil
}

// TODO подумать как улучшить
func PrepareAddParticipants(in usecases.AddParticipantsInput, partRepo repository.ParticipantRepository, boardRepo repository.BoardRepository) (usecases.AddParticipantsInput, error) {
	if _, err := boardRepo.GetById(in.BoardID); err != nil {
		return usecases.AddParticipantsInput{}, err
	}
	if _, err := partRepo.GetIdByTgUserId(in.BoardOwnerTgUserID); err != nil {
		return usecases.AddParticipantsInput{}, err
	}
	list, err := partRepo.GetIdByTgUserIdList(in.ParticipantTgUserIDs)
	if err != nil {
		return usecases.AddParticipantsInput{}, err
	}
	in.ParticipantsDB = list
	return in, nil
}

func PrepareGetBoards(tgUserId int64, repo repository.ParticipantRepository) (int64, error) {
	return repo.GetIdByTgUserId(tgUserId)
}
