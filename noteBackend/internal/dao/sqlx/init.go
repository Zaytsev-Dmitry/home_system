package sqlx

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"noteBackendApp/config"
	"os"
)

func CreateSqlxPort(config *config.AppConfig) *SqlxAuthPort {
	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.Database.Username,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.DataBaseName,
	)

	db, err := sqlx.Connect("postgres", dbURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return &SqlxAuthPort{Db: db}
}
