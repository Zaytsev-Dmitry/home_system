package category

import "github.com/jmoiron/sqlx"

type CategoryRepositorySqlx struct {
	db *sqlx.DB
}

func NewCategorySqlx(db *sqlx.DB) *CategoryRepositorySqlx {
	return &CategoryRepositorySqlx{db: db}
}
