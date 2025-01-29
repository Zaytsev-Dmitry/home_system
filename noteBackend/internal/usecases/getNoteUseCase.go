package usecases

import (
	noteInterface "noteBackendApp/internal/dao/interface"
	noteDomain "noteBackendApp/internal/domain"
)

type GetNoteUseCase struct {
	DAO noteInterface.NoteDao
}

func (byId *GetNoteUseCase) GetNoteByAccountId(id int) []noteDomain.TelegramAccount {
	return byId.DAO.GetNotesByAccountId(id)
}
