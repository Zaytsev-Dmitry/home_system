package boardprepare

import (
	"expensia/internal/app/ports/out/dao/repository"
	"expensia/internal/app/prepare"
	"fmt"
)

type GetBoardsPreparer struct {
	ParticipantRepo repository.ParticipantRepository
}

func RegisterGetBoardsPreparer(reg *prepare.PrepareRegistry, repo repository.ParticipantRepository) {
	prepare.RegisterPreparer(reg, "get_boards", GetBoardsPreparer{ParticipantRepo: repo})
}

func (g GetBoardsPreparer) Prepare(input interface{}) (interface{}, error) {
	idTyped, ok := input.(int64)
	if !ok {
		return nil, fmt.Errorf("invalid input type for GetBoardsPreparer")
	}
	return g.ParticipantRepo.GetIdByTgUserId(idTyped)
}
