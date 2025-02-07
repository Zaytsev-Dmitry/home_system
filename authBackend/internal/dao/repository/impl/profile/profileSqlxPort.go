package profile

import (
	authServerDomain "authServer/internal/domain"
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	INSERT           = "insert into profile (account_id, role, telegram_username) values($1, $2, $3) RETURNING id, account_id, role, telegram_username"
	SELECT_BY_ACC_ID = "select * from profile where account_id=($1) limit 1"
)

type SqlxProfilePort struct {
	Db *sqlx.DB
}

func CreateSqlxProfilePort(db *sqlx.DB) *SqlxProfilePort {
	return &SqlxProfilePort{Db: db}
}

func (p *SqlxProfilePort) CreateProfile(account authServerDomain.Account, tgUsername string) authServerDomain.Profile {
	var result authServerDomain.Profile
	rowx := p.Db.QueryRowx(INSERT, account.ID, "USER", tgUsername)
	if rowx.Err() != nil {

	} else {
		err := rowx.StructScan(&result)
		if err != nil {
			//TODO кинуть ошибку
		}
	}
	return result
}

func (p *SqlxProfilePort) GetProfileByAccountId(accId int64) authServerDomain.Profile {
	var resp authServerDomain.Profile
	err := p.Db.Get(&resp, SELECT_BY_ACC_ID, accId)
	if err != nil {
		//TODO кинуть ошибку и потом создать базовый профиль
		fmt.Println(err)
	}
	return resp
}

func (p *SqlxProfilePort) CloseConnection() {
	p.Db.Close()
}
