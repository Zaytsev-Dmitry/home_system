package dao

import (
	"expensia/configs"
	"expensia/internal/app/ports/out/dao/repository"
	"expensia/internal/infrastructure/persistence/board"
	"expensia/internal/infrastructure/persistence/boardParticipant"
	"expensia/internal/infrastructure/persistence/category"
	"expensia/internal/infrastructure/persistence/currency"
	"expensia/internal/infrastructure/persistence/expense"
	"expensia/internal/infrastructure/persistence/expenseShare"
	"expensia/internal/infrastructure/persistence/participant"
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
		BoardRepo:            board.NewBoardRepositorySqlx(db),
		BoardParticipantRepo: boardParticipant.NewBoardParticipantRepositorySqlx(db),
		CategoryRepo:         category.NewCategoryRepositorySqlx(db),
		CurrencyRepo:         currency.NewCurrencyRepositorySqlx(db),
		ExpenseRepo:          expense.NewExpenseRepositorySqlx(db),
		ExpenseShareRepo:     expenseShare.NewExpenseShareRepositorySqlx(db),
		ParticipantRepo:      participant.NewParticipantRepositorySqlx(db),
	}, db
}
