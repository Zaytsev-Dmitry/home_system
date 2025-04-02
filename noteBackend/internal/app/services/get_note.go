package services

import (
	"noteBackendApp/internal/dao"
	"noteBackendApp/internal/domain"
)

type GetNoteUCaseImpl struct {
	Dao dao.NoteDao
}

func (g *GetNoteUCaseImpl) GetNoteByTgId(id int64) []domain.Note {
	return g.Dao.GetNotesByTgId(id)
}
