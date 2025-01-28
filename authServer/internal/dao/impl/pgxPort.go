package authDaoPorts

import (
	authServerDomain "authServer/internal/domain"
	"github.com/jmoiron/sqlx"
)

type SqlxAuthPort struct {
	Db *sqlx.DB
}

func (port *SqlxAuthPort) Save(entity authServerDomain.Account) authServerDomain.Account {
	tx := port.Db.MustBegin()
	var account authServerDomain.Account
	err := port.Db.QueryRowx(
		"insert into accounts (first_name, last_name, login, email) values($1, $2, $3, $4) RETURNING id, first_name, last_name, login, email",
		entity.FirstName, entity.LastName, entity.Login, entity.Email).Scan(&account.ID, &account.FirstName, &account.LastName, &account.Login, &account.Email)
	if err != nil {
		//TODO кинуть ошибку
		tx.Rollback()
	}
	//TODO кинуть ошибку
	err = tx.Commit()
	if err != nil {
		return authServerDomain.Account{}
	}
	return account
}

func (port *SqlxAuthPort) CloseConnection() {
	port.Db.Close()
}

func CreateSqlxAuthPort(db *sqlx.DB) *SqlxAuthPort {
	return &SqlxAuthPort{Db: db}
}
