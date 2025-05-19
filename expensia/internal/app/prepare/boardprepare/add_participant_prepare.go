package boardprepare

import (
	"database/sql"
	"errors"
	"expensia/internal/app/ports/out/dao/repository"
	"expensia/internal/app/prepare"
	"fmt"
	"github.com/Zaytsev-Dmitry/apikit/custom_errors"
)

type AddParticipantPreparer struct {
	ParticipantRepo repository.ParticipantRepository
	BoardRepo       repository.BoardRepository
}

func RegisterAddParticipantPreparer(reg *prepare.PrepareRegistry, repo repository.ParticipantRepository, boardRepo repository.BoardRepository) {
	prepare.RegisterPreparer(reg, "add_participant_to_board", AddParticipantPreparer{
		ParticipantRepo: repo,
		BoardRepo:       boardRepo,
	})
}

func (a AddParticipantPreparer) Prepare(input interface{}) (interface{}, error) {
	req, ok := input.(repository.AddParticipantsInput)
	if !ok {
		return nil, fmt.Errorf("invalid input type for AddParticipantPreparer")
	}
	_, err := a.BoardRepo.GetById(req.BoardID)

	if err != nil {
		return nil, err
	}

	_, err = a.ParticipantRepo.GetIdByTgUserId(req.BoardOwnerTgUserID)
	if err != nil {
		return nil, err
	}

	list, err := a.ParticipantRepo.GetIdByTgUserIdList(req.ParticipantTgUserIDs)
	if err != nil && (errors.Is(err, sql.ErrNoRows) || len(list) == 0) {
		return nil, custom_errors.RowNotFound
	}
	req.ParticipantsDB = list
	return req, nil
}
