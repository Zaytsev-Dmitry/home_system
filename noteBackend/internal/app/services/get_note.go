package services

import (
	"noteBackendApp/internal/dao"
	"noteBackendApp/internal/domain"
)

type GetNoteUCase struct {
	Dao dao.NoteDao
}

func (g *GetNoteUCase) GetNoteByTgId(id int64) []domain.Note {
	return g.Dao.GetNotesByTgId(id)
}
