package dao

import (
	authConfig "authServer/configs"
	"authServer/internal/dao/repository/impl/account"
	implAccount "authServer/internal/dao/repository/impl/profile"
	"authServer/internal/dao/repository/intefraces"
	"fmt"
	"github.com/jmoiron/sqlx"
	"os"
)

type AuthDao struct {
	AccountRepo intefraces.AccountRepository
	ProfileRepo intefraces.ProfileRepository
}

func CreateDao(config authConfig.AppConfig) *AuthDao {
	var accRepo intefraces.AccountRepository
	var profileRepo intefraces.ProfileRepository

	initRepos(&accRepo, &profileRepo, &config)
	return &AuthDao{
		AccountRepo: accRepo,
		ProfileRepo: profileRepo,
	}
}

func initRepos(acc *intefraces.AccountRepository, prof *intefraces.ProfileRepository, config *authConfig.AppConfig) {
	if config.Database.Impl == "sqlx" {
		db := initSqlxDB(config)
		*acc = account.CreateSqlxAccountPort(db)
		*prof = implAccount.CreateSqlxProfilePort(db)
	}
}

func initSqlxDB(config *authConfig.AppConfig) *sqlx.DB {
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
