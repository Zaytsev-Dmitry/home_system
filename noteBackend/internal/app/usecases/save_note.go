package usecases

import (
	"noteBackendApp/internal/app/domain"
	"noteBackendApp/internal/app/ports/out/dao/dto"
)

type SaveNoteUCase interface {
	Save(dto dto.CreateNoteDto) (*domain.Note, error)
}
