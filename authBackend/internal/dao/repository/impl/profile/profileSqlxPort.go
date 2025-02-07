package profile

import (
	"authServer/internal/dao/repository"
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

func (p *SqlxProfilePort) CreateProfile(account authServerDomain.Account, tgUsername string) error {
	var result authServerDomain.Profile
	var resultErr error

	tx := p.Db.MustBegin()
	if p.GetProfileByAccountId(int64(account.ID)).ID == 0 {
		insertErr := tx.QueryRowx(INSERT, account.ID, "USER", tgUsername).StructScan(&result)
		resultErr = repository.ProceedInsertErrorsWithCallback(insertErr, tx)
	}
	resultErr = repository.CommitAndProceedErrors(tx, resultErr)

	return resultErr
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
