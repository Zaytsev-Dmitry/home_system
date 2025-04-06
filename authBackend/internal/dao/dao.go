package dao

import (
	"authServer/internal/dao/repository/impl/account"
	"authServer/internal/dao/repository/impl/profile"
	"authServer/internal/dao/repository/intefraces"
	"authServer/pkg/config_loader"
	"fmt"
	"github.com/jmoiron/sqlx"
	"os"
)

type AuthDao struct {
	AccountRepo intefraces.AccountRepository
	ProfileRepo intefraces.ProfileRepository
}

func New(config *config_loader.AppConfig) *AuthDao {
	var accRepo intefraces.AccountRepository
	var profileRepo intefraces.ProfileRepository

	if config.Database.Impl == "sqlx" {
		db := newSqlxDB(config)
		accRepo = account.New(db)
		profileRepo = profile.New(db)
	}

	return &AuthDao{
		AccountRepo: accRepo,
		ProfileRepo: profileRepo,
	}
}

func newSqlxDB(config *authConfig.AppConfig) *sqlx.DB {
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
