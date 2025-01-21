package usecases

import (
	noteDao "noteBackendApp/internal/dao"
	noteDomain "noteBackendApp/internal/domain"
)

type SaveNoteUseCase struct {
	Db *noteDao.InMemoryNoteRepository
}

func (saveUse *SaveNoteUseCase) Save(toSave noteDomain.NoteEntity) noteDomain.NoteEntity {
	found, err := saveUse.Db.GetById(toSave.Id)
	if err != nil {
		return saveUse.Db.Save(toSave)
	} else {
		return found
	}
}
