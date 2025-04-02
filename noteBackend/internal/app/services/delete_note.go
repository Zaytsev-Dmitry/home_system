package services

import (
	"noteBackendApp/internal/app/ports/out/dao/repository/note"
	customErr "noteBackendApp/pkg/errors"
)

type DeleteNoteUCaseImpl struct {
	NoteRepo note.NoteRepository
}

func (d *DeleteNoteUCaseImpl) DeleteNoteByTgId(tgId int64) error {
	err := d.NoteRepo.DeleteNotesByTgId(tgId)

	if err != nil {
		return customErr.UpdateErrorText(err.Action, err.WrapErr)
	}
	return nil
}
