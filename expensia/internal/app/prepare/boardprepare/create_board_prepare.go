package boardprepare

import (
	"expensia/internal/app/ports/out/dao/repository"
	"expensia/internal/app/prepare"
	"fmt"
)

type CreateBoardPreparer struct {
	ParticipantRepo repository.ParticipantRepository
}

func RegisterCreateBoardPreparer(reg *prepare.PrepareRegistry, repo repository.ParticipantRepository) {
	prepare.RegisterPreparer(reg, "create_board", CreateBoardPreparer{ParticipantRepo: repo})
}

func (p CreateBoardPreparer) Prepare(input interface{}) (interface{}, error) {
	req, ok := input.(repository.CreateBoardUCaseIn)
	if !ok {
		return nil, fmt.Errorf("invalid input type for CreateBoardPreparer")
	}
	id, err := p.ParticipantRepo.GetIdByTgUserId(req.TgUserId)
	if err != nil {
		return repository.CreateBoardUCaseIn{}, err
	}
	req.OwnerId = id
	return req, nil
}
