package services

import (
	"noteBackendApp/internal/dao"
	"noteBackendApp/internal/domain"
)

type SaveNoteUCase struct {
	Dao dao.NoteDao
}

func (s *SaveNoteUCase) Save(toSave domain.Note) domain.Note {
	found := s.Dao.ExistByName(toSave.Name)
	if found {
		return s.Dao.GetByName(toSave.Name)
	} else {
		return s.Dao.Save(toSave)
	}
}
