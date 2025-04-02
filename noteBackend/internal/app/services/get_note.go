package services

import (
	"noteBackendApp/internal/app/domain"
	"noteBackendApp/internal/dao"
)

type GetNoteUCaseImpl struct {
	Dao dao.NoteDao
}

func (g *GetNoteUCaseImpl) GetNoteByTgId(id int64) []domain.Note {
	return g.Dao.GetNotesByTgId(id)
}
