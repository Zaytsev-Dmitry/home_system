package expenseShare

import "github.com/jmoiron/sqlx"

type ExpenseShareRepositorySqlx struct {
	db *sqlx.DB
}

func NewExpenseShareSqlx(db *sqlx.DB) *ExpenseShareRepositorySqlx {
	return &ExpenseShareRepositorySqlx{db: db}
}
