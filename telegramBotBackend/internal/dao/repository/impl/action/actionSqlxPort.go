package action

import (
	"github.com/jmoiron/sqlx"
)

const (
	INSERT_PROFILE   = "insert into profile (account_id, role, is_active, username) values($1, $2, $3, $4) RETURNING id, account_id, role, is_active, username"
	SELECT_BY_ACC_ID = "select * from profile where account_id=($1) limit 1"
)

type SqlxActionPort struct {
	Db *sqlx.DB
}

func CreateSqlxActionPort(db *sqlx.DB) *SqlxActionPort {
	return &SqlxActionPort{Db: db}
}

func (p *SqlxActionPort) CloseConnection() {
	p.Db.Close()
}
