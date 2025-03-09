package dao

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"os"
	"telegramCLient/config"
	"telegramCLient/internal/dao/repository/impl/action"
	"telegramCLient/internal/dao/repository/intefraces"
)

type TelegramBotDao struct {
	ActionRepo intefraces.ActionRepository
}

func CreateDao(config config.AppConfig) *TelegramBotDao {
	var actionRepo intefraces.ActionRepository
	initRepos(&actionRepo, &config)
	return &TelegramBotDao{
		ActionRepo: actionRepo,
	}
}

func initRepos(acc *intefraces.ActionRepository, config *config.AppConfig) {
	if config.Database.Impl == "sqlx" {
		db := initSqlxDB(config)
		*acc = action.CreateSqlxActionPort(db)
	}
}

func initSqlxDB(config *config.AppConfig) *sqlx.DB {
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

func (dao *TelegramBotDao) Close() {
	dao.ActionRepo.CloseConnection()
}
