package expense

import "github.com/jmoiron/sqlx"

type ExpenseRepositorySqlx struct {
	db *sqlx.DB
}

func NewExpenseSqlx(db *sqlx.DB) *ExpenseRepositorySqlx {
	return &ExpenseRepositorySqlx{db: db}
}
