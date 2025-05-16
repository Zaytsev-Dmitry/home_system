package expense

import "github.com/jmoiron/sqlx"

type ExpenseRepositorySqlx struct {
	db *sqlx.DB
}

func NewExpenseRepositorySqlx(db *sqlx.DB) *ExpenseRepositorySqlx {
	return &ExpenseRepositorySqlx{db: db}
}
