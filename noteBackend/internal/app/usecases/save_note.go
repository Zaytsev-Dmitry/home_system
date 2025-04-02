package usecases

import (
	"noteBackendApp/internal/app/domain"
)

type SaveNoteUCase interface {
	Save(toSave domain.Note) domain.Note
}
