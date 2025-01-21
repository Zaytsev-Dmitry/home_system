package usecases

import (
	noteDao "noteBackendApp/internal/dao"
	noteDomain "noteBackendApp/internal/domain"
)

type GetNoteUseCase struct {
	Db *noteDao.InMemoryNoteRepository
}

func (byId *GetNoteUseCase) GetById(id string) (noteDomain.NoteEntity, error) {
	return byId.Db.GetById(id)
}

func (all *GetNoteUseCase) GetAll() []noteDomain.NoteEntity {
	return all.GetAll()
}
