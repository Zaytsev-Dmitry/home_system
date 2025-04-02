package services

import (
	"noteBackendApp/internal/app/domain"
	"noteBackendApp/internal/app/ports/out/dao/repository/note"
)

type SaveNoteUCaseImpl struct {
	NoteRepo note.NoteRepository
}

func (s *SaveNoteUCaseImpl) Save(toSave domain.Note) domain.Note {
	found := s.NoteRepo.ExistByName(toSave.Name)
	if found {
		return s.NoteRepo.GetByName(toSave.Name)
	} else {
		return s.NoteRepo.Save(toSave)
	}
}
