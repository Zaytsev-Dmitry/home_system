package dao

import (
	intefraces2 "authServer/internal/app/ports/out/dao/repository/intefraces"
	"authServer/pkg/config_loader"
	"fmt"
	"github.com/jmoiron/sqlx"
	"os"
)

type AuthDao struct {
	AccountRepo intefraces2.AccountRepository
	ProfileRepo intefraces2.ProfileRepository
}

func New(config *config_loader.AppConfig) *AuthDao {
	var accRepo intefraces2.AccountRepository
	var profileRepo intefraces2.ProfileRepository

	return &AuthDao{
		AccountRepo: accRepo,
		ProfileRepo: profileRepo,
	}
}

func newSqlxDB(config *config_loader.AppConfig) *sqlx.DB {
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

func (dao *AuthDao) Close() {
	dao.AccountRepo.CloseConnection()
	dao.ProfileRepo.CloseConnection()
}
