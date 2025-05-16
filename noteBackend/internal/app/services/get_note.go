package services

import (
	"noteBackendApp/internal/app/domain"
	"noteBackendApp/internal/app/ports/out/dao/repository/note"
	customErr "noteBackendApp/pkg/errors"
)

type GetNoteUCaseImpl struct {
	NoteRepo note.NoteRepository
}

func (g *GetNoteUCaseImpl) GetNoteByTgId(id int64) ([]*domain.Note, error) {
	objects, err := g.NoteRepo.GetNotesByTgId(id)

	if err != nil {
		return nil, customErr.UpdateErrorText(err.Action, err.WrapErr)
	}

	return objects, nil
}
