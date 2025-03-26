package usecases

import (
	"noteBackendApp/internal/dao"
	"noteBackendApp/internal/domain"
)

type GetNoteUseCase struct {
	Dao dao.NoteDao
}

func (byId *GetNoteUseCase) GetNoteByTgId(id int64) []domain.Note {
	return byId.Dao.GetNotesByTgId(id)
}
