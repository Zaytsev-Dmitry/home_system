package usecases

import (
	noteInterface "noteBackendApp/internal/dao/interface"
)

type DeleteNoteUseCase struct {
	DAO noteInterface.NoteDao
}

func (deleteUse *DeleteNoteUseCase) DeleteNoteByAccountId(id int) {
	deleteUse.DAO.DeleteNotesByAccountId(id)
}
