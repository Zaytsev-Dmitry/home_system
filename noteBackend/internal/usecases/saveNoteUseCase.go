package usecases

import (
	noteInterface "noteBackendApp/internal/dao/interface"
	noteDomain "noteBackendApp/internal/domain"
)

type SaveNoteUseCase struct {
	DAO noteInterface.NoteDao
}

func (saveUse *SaveNoteUseCase) Save(toSave noteDomain.TelegramAccount) noteDomain.TelegramAccount {
	found := saveUse.DAO.ExistByName(toSave.Name)
	if found {
		return saveUse.DAO.GetByName(toSave.Name)
	} else {
		return saveUse.DAO.Save(toSave)
	}
}
