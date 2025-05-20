package persistence

import "github.com/jmoiron/sqlx"

type CurrencyRepositorySqlx struct {
	db *sqlx.DB
}

func NewCurrencyRepositorySqlx(db *sqlx.DB) *CurrencyRepositorySqlx {
	return &CurrencyRepositorySqlx{db: db}
}
