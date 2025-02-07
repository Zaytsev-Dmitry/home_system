package account

import (
	authServerDomain "authServer/internal/domain"
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	INSERT_ACCOUNT     = "insert into accounts (first_name, last_name, email, type, telegram_id) values($1, $2, $3, $4, $5) RETURNING id, first_name, last_name, email, type"
	SELECT_ID_BY_TG_ID = "select ac.id from accounts ac where ac.telegram_id = ($1) limit 1"
)

type SqlxAccountPort struct {
	Db *sqlx.DB
}

func (port *SqlxAccountPort) Save(entity authServerDomain.Account) (authServerDomain.Account, error) {
	tx := port.Db.MustBegin()
	var account authServerDomain.Account
	err := port.Db.QueryRowx(INSERT_ACCOUNT, entity.FirstName, entity.LastName, entity.Email, entity.Type, entity.TelegramId).StructScan(&account)
	if err != nil {
		//TODO кинуть ошибку
		tx.Rollback()
	}
	//TODO кинуть ошибку
	err = tx.Commit()
	if err != nil {
		return authServerDomain.Account{}, nil
	}
	return account, nil
}

func (port *SqlxAccountPort) GetIdByTgId(tgId int64) int64 {
	var resp int64
	err := port.Db.Get(&resp, SELECT_ID_BY_TG_ID, tgId)
	if err != nil {
		//TODO кинуть ошибку и потом создать базовый профиль
		fmt.Println(err)
	}
	return resp
}

func (port *SqlxAccountPort) CloseConnection() {
	port.Db.Close()
}

func CreateSqlxAccountPort(db *sqlx.DB) *SqlxAccountPort {
	return &SqlxAccountPort{Db: db}
}
