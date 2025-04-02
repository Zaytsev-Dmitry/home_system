package usecases

import (
	"noteBackendApp/internal/app/domain"
)

type GetNoteUCase interface {
	GetNoteByTgId(id int64) []domain.Note
}
