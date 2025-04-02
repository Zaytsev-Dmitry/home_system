package note

import (
	"noteBackendApp/internal/app/domain"
	"noteBackendApp/pkg/errors"
)

type NoteRepository interface {
	Save(entity domain.Note) (*domain.Note, *errors.CustomError)
	DeleteNotesByTgId(tgId int64) *errors.CustomError
	GetNotesByTgId(tgId int64) ([]*domain.Note, *errors.CustomError)
	ExistByName(name string) (*domain.Note, *errors.CustomError)
}
