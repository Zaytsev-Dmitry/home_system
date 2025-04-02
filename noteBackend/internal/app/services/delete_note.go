package services

import (
	"noteBackendApp/internal/app/ports/out/dao"
)

type DeleteNoteUCaseImpl struct {
	Dao dao.NoteDao
}

func (d *DeleteNoteUCaseImpl) DeleteNoteByTgId(tgId int64) {
	d.Dao.DeleteNotesByTgId(tgId)
}
