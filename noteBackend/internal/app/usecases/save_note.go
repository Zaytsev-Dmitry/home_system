package usecases

import "noteBackendApp/internal/domain"

type SaveNoteUCase interface {
	Save(toSave domain.Note) domain.Note
}
