package services

import (
	"noteBackendApp/internal/app/ports/out/dao/repository/note"
)

type DeleteNoteUCaseImpl struct {
	NoteRepo note.NoteRepository
}

func (d *DeleteNoteUCaseImpl) DeleteNoteByTgId(tgId int64) {
	d.NoteRepo.DeleteNotesByTgId(tgId)
}
