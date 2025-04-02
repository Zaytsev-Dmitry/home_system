package services

import "noteBackendApp/internal/dao"

type DeleteNoteUCase struct {
	Dao dao.NoteDao
}

func (d *DeleteNoteUCase) DeleteNoteByTgId(tgId int64) {
	d.Dao.DeleteNotesByTgId(tgId)
}
