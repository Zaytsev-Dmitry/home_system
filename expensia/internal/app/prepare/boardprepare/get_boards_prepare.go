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
	return g.ParticipantRepo.GetIdByTgUserId(input)
}
