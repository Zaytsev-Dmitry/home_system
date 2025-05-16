package boardprepare

import (
	"expensia/api/rest"
	"expensia/internal/app/ports/out/dao/repository"
	"expensia/internal/app/prepare"
)

type CreateBoardPreparer struct {
	ParticipantRepo repository.ParticipantRepository
}

func RegisterBoardPreparer(reg *prepare.PrepareRegistry, repo repository.ParticipantRepository) {
	prepare.RegisterPreparer(reg, "create_board", CreateBoardPreparer{ParticipantRepo: repo})
}

func (p CreateBoardPreparer) Prepare(input rest.CreateBoardParams) (repository.CreateBoardUCaseIn, error) {
	id, err := p.ParticipantRepo.GetIdByTgUserId(input.TgUserId)
	if err != nil {
		return repository.CreateBoardUCaseIn{}, err
	}
	return repository.CreateBoardUCaseIn{
		OwnerId:  id,
		Name:     input.Name,
		Currency: string(input.Currency),
	}, nil
}
