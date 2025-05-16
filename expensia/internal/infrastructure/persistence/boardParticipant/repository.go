package boardParticipant

import "github.com/jmoiron/sqlx"

type BoardParticipantRepositorySqlx struct {
	db *sqlx.DB
}

func NewBoardParticipantRepositorySqlx(db *sqlx.DB) *BoardParticipantRepositorySqlx {
	return &BoardParticipantRepositorySqlx{db: db}
}
