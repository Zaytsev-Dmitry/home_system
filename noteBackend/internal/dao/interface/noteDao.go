package noteDaoInterface

import (
	noteDomain "noteBackendApp/internal/domain"
)

type NoteDao interface {
	Save(entity noteDomain.Note) noteDomain.Note
	DeleteNotesByTgId(tgId int64)
	GetNotesByTgId(tgId int64) []noteDomain.Note
	ExistByName(name string) bool
	GetByName(name string) noteDomain.Note
}
