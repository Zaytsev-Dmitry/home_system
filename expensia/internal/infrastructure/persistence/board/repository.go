package board

import "github.com/jmoiron/sqlx"

type BoardRepositorySqlx struct {
	db *sqlx.DB
}

func NewBoardRepositorySqlx(db *sqlx.DB) *BoardRepositorySqlx {
	return &BoardRepositorySqlx{db: db}
}
