package usecases

import (
	noteInterface "noteBackendApp/internal/dao/interface"
	noteDomain "noteBackendApp/internal/domain"
)

type GetNoteUseCase struct {
	DAO noteInterface.NoteDao
}

func (byId *GetNoteUseCase) GetNoteByTgId(id int64) []noteDomain.Note {
	return byId.DAO.GetNotesByTgId(id)
}
