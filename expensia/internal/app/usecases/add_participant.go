package usecases

import "expensia/internal/app/ports/out/dao/repository"

type AddParticipantUCase interface {
	AddParticipantsToBoard(req repository.AddParticipantsInput) error
}
