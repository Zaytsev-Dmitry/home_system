package dao

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"os"
	"userService/internal/app/ports/out/dao/repository/account"
	"userService/internal/app/ports/out/dao/repository/profile"
	"userService/pkg/config_loader"
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
