package services

import (
	"noteBackendApp/internal/app/domain"
	"noteBackendApp/internal/app/ports/out/dao"
)

type GetNoteUCaseImpl struct {
	Dao dao.NoteDao
}

func (g *GetNoteUCaseImpl) GetNoteByTgId(id int64) []domain.Note {
	return g.Dao.NoteRepo.GetNotesByTgId(id)
}
