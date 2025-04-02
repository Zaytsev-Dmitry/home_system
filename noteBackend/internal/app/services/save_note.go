package services

import (
	"noteBackendApp/internal/app/domain"
	"noteBackendApp/internal/app/ports/out/dao"
)

type SaveNoteUCaseImpl struct {
	Dao *dao.NoteDao
}

func (s *SaveNoteUCaseImpl) Save(toSave domain.Note) domain.Note {
	found := s.Dao.NoteRepo.ExistByName(toSave.Name)
	if found {
		return s.Dao.NoteRepo.GetByName(toSave.Name)
	} else {
		return s.Dao.NoteRepo.Save(toSave)
	}
}
