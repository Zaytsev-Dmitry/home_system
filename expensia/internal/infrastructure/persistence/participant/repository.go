package participant

import "github.com/jmoiron/sqlx"

type ParticipantRepositorySqlx struct {
	db *sqlx.DB
}

func NewParticipantRepositorySqlx(db *sqlx.DB) *ParticipantRepositorySqlx {
	return &ParticipantRepositorySqlx{db: db}
}
