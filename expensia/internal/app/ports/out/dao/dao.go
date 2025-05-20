package dao

import (
	"expensia/configs"
	"expensia/internal/app/ports/out/dao/repository"
	"expensia/internal/infrastructure/persistence"
	"fmt"
	"github.com/jmoiron/sqlx"
	"os"
)

type ExpensiaDao struct {
	BoardRepo            repository.BoardRepository
	BoardParticipantRepo repository.BoardParticipantRepository
	CategoryRepo         repository.CategoryRepository
	CurrencyRepo         repository.CurrencyRepository
	ExpenseRepo          repository.ExpenseRepository
	ExpenseShareRepo     repository.ExpenseShareRepository
	ParticipantRepo      repository.ParticipantRepository
}

func newDbConnection(config *configs.AppConfig) *sqlx.DB {
	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:5432/%s?sslmode=disable",
		config.Database.Username,
		config.Database.Password,
		config.Database.Host,
		config.Database.DataBaseName,
	)
	db, err := sqlx.Connect("postgres", dbURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return db
}

func Create(config *configs.AppConfig) (*ExpensiaDao, *sqlx.DB) {
	db := newDbConnection(config)
	return &ExpensiaDao{
		BoardRepo:            persistence.NewBoardRepositorySqlx(db),
		BoardParticipantRepo: persistence.NewBoardParticipantRepositorySqlx(db),
		CategoryRepo:         persistence.NewCategoryRepositorySqlx(db),
		CurrencyRepo:         persistence.NewCurrencyRepositorySqlx(db),
		ExpenseRepo:          persistence.NewExpenseRepositorySqlx(db),
		ExpenseShareRepo:     persistence.NewExpenseShareRepositorySqlx(db),
		ParticipantRepo:      persistence.NewParticipantRepositorySqlx(db),
	}, db
}
