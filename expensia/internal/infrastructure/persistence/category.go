package persistence

import "github.com/jmoiron/sqlx"

type CategoryRepositorySqlx struct {
	db *sqlx.DB
}

func NewCategoryRepositorySqlx(db *sqlx.DB) *CategoryRepositorySqlx {
	return &CategoryRepositorySqlx{db: db}
}
