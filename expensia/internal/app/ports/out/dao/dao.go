package dao

import (
	"expensia/configs"
	"expensia/internal/app/ports/out/dao/repository/board"
	"expensia/internal/app/ports/out/dao/repository/boardParticipant"
	"expensia/internal/app/ports/out/dao/repository/category"
	"expensia/internal/app/ports/out/dao/repository/currency"
	"expensia/internal/app/ports/out/dao/repository/expense"
	"expensia/internal/app/ports/out/dao/repository/expenseShare"
	"expensia/internal/app/ports/out/dao/repository/participant"
	"fmt"
	"github.com/jmoiron/sqlx"
	"os"
)

type ExpensiaDao struct {
	BoardRepo            board.BoardRepository
	BoardParticipantRepo boardParticipant.BoardParticipantRepository
	CategoryRepo         category.CategoryRepository
	CurrencyRepo         currency.CurrencyRepository
	ExpenseRepo          expense.ExpenseRepository
	ExpenseShareRepo     expenseShare.ExpenseShareRepository
	ParticipantRepo      participant.ParticipantRepository
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
		BoardRepo:            board.NewBoardSqlx(db),
		BoardParticipantRepo: boardParticipant.NewBoardParticipantSqlx(db),
		CategoryRepo:         category.NewCategorySqlx(db),
		CurrencyRepo:         currency.NewCurrencySqlx(db),
		ExpenseRepo:          expense.NewExpenseSqlx(db),
		ExpenseShareRepo:     expenseShare.NewExpenseShareSqlx(db),
		ParticipantRepo:      participant.NewParticipantSqlx(db),
	}, db
}
