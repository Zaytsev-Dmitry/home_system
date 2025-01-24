package usecases

import (
	noteDao "noteBackendApp/internal/dao/impl"
	noteDomain "noteBackendApp/internal/domain"
)

type SaveNoteUseCase struct {
	Db *noteDao.InMemoryNoteRepository
}

func (saveUse *SaveNoteUseCase) Save(toSave noteDomain.TelegramAccount) noteDomain.TelegramAccount {
	found, err := saveUse.Db.GetNoteByAccountId(toSave.AccountId)
	if err != nil {
		return saveUse.Db.Save(toSave)
	} else {
		return found
	}
}
