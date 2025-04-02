package usecases

import "noteBackendApp/internal/domain"

type GetNoteUCase interface {
	GetNoteByTgId(id int64) []domain.Note
}
