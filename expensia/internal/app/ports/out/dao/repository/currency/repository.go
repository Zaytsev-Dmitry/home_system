package currency

import "github.com/jmoiron/sqlx"

type CurrencyRepositorySqlx struct {
	db *sqlx.DB
}

func NewCurrencySqlx(db *sqlx.DB) *CurrencyRepositorySqlx {
	return &CurrencyRepositorySqlx{db: db}
}
