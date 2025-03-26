package usecases

import (
	"noteBackendApp/internal/dao"
	"noteBackendApp/internal/domain"
)

type SaveNoteUseCase struct {
	Dao dao.NoteDao
}

func (saveUse *SaveNoteUseCase) Save(toSave domain.Note) domain.Note {
	found := saveUse.Dao.ExistByName(toSave.Name)
	if found {
		return saveUse.Dao.GetByName(toSave.Name)
	} else {
		return saveUse.Dao.Save(toSave)
	}
}
