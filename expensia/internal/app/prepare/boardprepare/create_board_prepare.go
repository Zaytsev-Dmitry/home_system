package boardprepare

import (
	"expensia/internal/app/ports/out/dao/repository"
	"expensia/internal/app/prepare"
)

type CreateBoardPreparer struct {
	ParticipantRepo repository.ParticipantRepository
}

func (p CreateBoardPreparer) Prepare(input repository.CreateBoardUCaseIn) (repository.CreateBoardUCaseIn, error) {
	id, err := p.ParticipantRepo.GetIdByTgUserId(input.TgUserId)
	if err != nil {
		return repository.CreateBoardUCaseIn{}, err
	}
	input.OwnerId = id
	return input, nil
}

func RegisterCreateBoardPreparer(reg *prepare.PrepareRegistry, repo repository.ParticipantRepository) {
	prepare.RegisterPreparer(reg, "create_board", CreateBoardPreparer{ParticipantRepo: repo})
}
