package usecases

import (
	noteDao "noteBackendApp/internal/dao"
	noteDomain "noteBackendApp/internal/domain"
)

type GetNoteUseCase struct {
	Db *noteDao.InMemoryNoteRepository
}

func (byId *GetNoteUseCase) GetNoteByAccountId(id int) (noteDomain.NoteEntity, error) {
	return byId.Db.GetNoteByAccountId(id)
}
