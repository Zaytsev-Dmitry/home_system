package boardprepare

import (
	"expensia/internal/app/ports/out/dao/repository"
	"expensia/internal/app/prepare"
)

type GetBoardsPreparer struct {
	ParticipantRepo repository.ParticipantRepository
}

func RegisterGetBoardsPreparer(reg *prepare.PrepareRegistry, repo repository.ParticipantRepository) {
	prepare.RegisterPreparer(reg, "get_boards", GetBoardsPreparer{ParticipantRepo: repo})
}

func (g GetBoardsPreparer) Prepare(input int64) (int64, error) {
	id, err := g.ParticipantRepo.GetIdByTgUserId(input)
	if err != nil {
		return 0, err
	}
	return id, nil
}
