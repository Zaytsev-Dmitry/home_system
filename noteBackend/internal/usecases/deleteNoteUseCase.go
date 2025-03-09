package usecases

import (
	noteInterface "noteBackendApp/internal/dao/interface"
)

type DeleteNoteUseCase struct {
	DAO noteInterface.NoteDao
}

func (deleteUse *DeleteNoteUseCase) DeleteNoteByTgId(tgId int64) {
	deleteUse.DAO.DeleteNotesByTgId(tgId)
}
