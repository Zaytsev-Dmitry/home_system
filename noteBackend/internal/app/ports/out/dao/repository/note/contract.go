package note

import (
	"noteBackendApp/internal/app/domain"
)

type NoteRepository interface {
	Save(entity domain.Note) domain.Note
	DeleteNotesByTgId(tgId int64)
	GetNotesByTgId(tgId int64) []domain.Note
	ExistByName(name string) bool
	GetByName(name string) domain.Note
}
