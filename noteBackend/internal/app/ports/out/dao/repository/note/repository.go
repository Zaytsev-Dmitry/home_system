package note

import (
	"github.com/jmoiron/sqlx"
	"noteBackendApp/internal/app/domain"
)

const (
	Get       = "Get"
	QueryRowx = "QueryRowx"
)

type NoteRepositorySqlx struct {
	db *sqlx.DB
}

func (n NoteRepositorySqlx) Save(entity domain.Note) domain.Note {
	return domain.Note{}
}

func (n NoteRepositorySqlx) DeleteNotesByTgId(tgId int64) {
	return
}

func (n NoteRepositorySqlx) GetNotesByTgId(tgId int64) []domain.Note {
	return []domain.Note{}
}

func (n NoteRepositorySqlx) ExistByName(name string) bool {
	return false
}

func (n NoteRepositorySqlx) GetByName(name string) domain.Note {
	return domain.Note{}
}

func NewNoteSqlx(db *sqlx.DB) *NoteRepositorySqlx {
	return &NoteRepositorySqlx{db: db}
}
