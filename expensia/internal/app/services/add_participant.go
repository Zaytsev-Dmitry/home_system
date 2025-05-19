package services

import "expensia/internal/app/ports/out/dao/repository"

type AddParticipantUCaseImpl struct {
	Repo repository.BoardParticipantRepository
}

func (ap AddParticipantUCaseImpl) AddParticipantsToBoard(req repository.AddParticipantsInput) error {
	err := ap.Repo.AddParticipantsToBoard(
		req.BoardID,
		req.ParticipantsDB,
	)
	return err
}
