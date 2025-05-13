package participant

import "github.com/jmoiron/sqlx"

type ParticipantRepositorySqlx struct {
	db *sqlx.DB
}

func NewParticipantSqlx(db *sqlx.DB) *ParticipantRepositorySqlx {
	return &ParticipantRepositorySqlx{db: db}
}
