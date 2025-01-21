package usecases

import (
	noteDao "noteBackendApp/internal/dao"
)

type DeleteNoteUseCase struct {
	Db *noteDao.InMemoryNoteRepository
}

func (deleteUse *DeleteNoteUseCase) DeleteById(id string) {
	deleteUse.Db.DeleteById(id)
}
