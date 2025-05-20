package boardprepare

import (
	"expensia/internal/app/ports/out/dao/repository"
	"expensia/internal/app/prepare"
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

func (a AddParticipantPreparer) Prepare(req repository.AddParticipantsInput) (repository.AddParticipantsInput, error) {
	err := ReturnFirstError(
		func() error { _, err := a.BoardRepo.GetById(req.BoardID); return err },
		func() error { _, err := a.ParticipantRepo.GetIdByTgUserId(req.BoardOwnerTgUserID); return err },
		func() error {
			list, err := a.ParticipantRepo.GetIdByTgUserIdList(req.ParticipantTgUserIDs)
			if err == nil {
				req.ParticipantsDB = list
			}
			return err
		},
	)

	if err != nil {
		return repository.AddParticipantsInput{}, err
	}
	return req, nil
}

func ReturnFirstError(funcs ...func() error) error {
	for _, fn := range funcs {
		if err := fn(); err != nil {
			return err
		}
	}
	return nil
}
