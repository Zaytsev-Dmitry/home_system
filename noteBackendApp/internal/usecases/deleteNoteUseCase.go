package usecases

import (
	noteDao "noteBackendApp/internal/dao/impl"
)

type DeleteNoteUseCase struct {
	Db *noteDao.InMemoryNoteRepository
}

func (deleteUse *DeleteNoteUseCase) DeleteNoteByAccountId(id int) {
	deleteUse.Db.DeleteNoteByAccountId(id)
}
