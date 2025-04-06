package dao

import (
	"authBackend/internal/app/ports/out/dao/repository/account"
	"authBackend/internal/app/ports/out/dao/repository/profile"
	"authBackend/pkg/config_loader"
	"fmt"
	"github.com/jmoiron/sqlx"
	"os"
)

type AuthDao struct {
	AccountRepository account.AccountRepository
	ProfileRepository profile.ProfileRepository
}

func newDbConnection(config *config_loader.AppConfig) *sqlx.DB {
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

func Create(config *config_loader.AppConfig) (*AuthDao, *sqlx.DB) {
	db := newDbConnection(config)
	return &AuthDao{
		AccountRepository: account.NewAccountRepositorySqlx(db),
		ProfileRepository: profile.NewProfileRepositorySqlx(db),
	}, db
}
