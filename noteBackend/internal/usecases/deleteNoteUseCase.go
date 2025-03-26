package usecases

import (
	"noteBackendApp/internal/dao"
)

type DeleteNoteUseCase struct {
	Dao dao.NoteDao
}

func (deleteUse *DeleteNoteUseCase) DeleteNoteByTgId(tgId int64) {
	deleteUse.Dao.DeleteNotesByTgId(tgId)
}
