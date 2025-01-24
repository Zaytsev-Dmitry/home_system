package usecases

import (
	noteDao "noteBackendApp/internal/dao/impl"
	noteDomain "noteBackendApp/internal/domain"
)

type GetNoteUseCase struct {
	Db *noteDao.InMemoryNoteRepository
}

func (byId *GetNoteUseCase) GetNoteByAccountId(id int) (noteDomain.TelegramAccount, error) {
	return byId.Db.GetNoteByAccountId(id)
}
