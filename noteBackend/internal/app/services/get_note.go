package services

import (
	"noteBackendApp/internal/app/domain"
	"noteBackendApp/internal/app/ports/out/dao/repository/note"
)

type GetNoteUCaseImpl struct {
	NoteRepo note.NoteRepository
}

func (g *GetNoteUCaseImpl) GetNoteByTgId(id int64) []domain.Note {
	return g.NoteRepo.GetNotesByTgId(id)
}
