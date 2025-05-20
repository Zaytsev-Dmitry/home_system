package services

import (
	"expensia/internal/app/ports/out/dao/repository"
	"expensia/internal/app/usecases"
)

type AddParticipantUCaseImpl struct {
	Repo repository.BoardParticipantRepository
}

func (ap AddParticipantUCaseImpl) AddParticipantsToBoard(req usecases.AddParticipantsInput) error {
	err := ap.Repo.AddParticipantsToBoard(
		req.BoardID,
		req.ParticipantsDB,
	)
	return err
}
