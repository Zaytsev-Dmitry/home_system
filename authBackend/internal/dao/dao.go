package dao

import (
	authConfig "authServer/configs"
	"authServer/internal/dao/repository/account"
	implAccount "authServer/internal/dao/repository/account/impl"
	"authServer/internal/dao/repository/profile"
	implProfile "authServer/internal/dao/repository/profile/impl"
	"fmt"
	"github.com/jmoiron/sqlx"
	"os"
)

type AuthDao struct {
	AccountRepo account.AccountRepository
	ProfileRepo profile.ProfileRepository
}

func CreateDao(config authConfig.AppConfig) *AuthDao {
	var accRepo account.AccountRepository
	var profileRepo profile.ProfileRepository

	initRepos(&accRepo, &profileRepo, &config)
	return &AuthDao{
		AccountRepo: accRepo,
		ProfileRepo: profileRepo,
	}
}

func initRepos(acc *account.AccountRepository, prof *profile.ProfileRepository, config *authConfig.AppConfig) {
	if config.Database.Impl == "sqlx" {
		db := initSqlxDB(config)
		*acc = implAccount.CreateSqlxAccountPort(db)
		*prof = implProfile.CreateSqlxProfilePort(db)
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
