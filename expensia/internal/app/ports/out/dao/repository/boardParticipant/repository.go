package boardParticipant

import (
	"github.com/jmoiron/sqlx"
)

type BoardParticipantRepositorySqlx struct {
	db *sqlx.DB
}

func NewBoardParticipantSqlx(db *sqlx.DB) *BoardParticipantRepositorySqlx {
	return &BoardParticipantRepositorySqlx{db: db}
}
