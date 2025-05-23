package dao

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"os"
	"userService/configs"
	"userService/internal/app/ports/out/dao/repository/identity"
	"userService/internal/app/ports/out/dao/repository/profile"
)

type UserDao struct {
	IdentityRepo      identity.UserIdentityLinkRepository
	ProfileRepository profile.ProfileRepository
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

func Create(config *configs.AppConfig) (*UserDao, *sqlx.DB) {
	db := newDbConnection(config)
	return &UserDao{
		IdentityRepo:      identity.NewUserIdentityLinkRepositorySqlx(db),
		ProfileRepository: profile.NewProfileRepositorySqlx(db),
	}, db
}
