package account

import (
	"authServer/internal/dao/repository"
	"authServer/internal/dao/repository/intefraces"
	authServerDomain "authServer/internal/domain"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	INSERT_ACCOUNT     = "insert into accounts (first_name, last_name, email, type, telegram_id) values($1, $2, $3, $4, $5) RETURNING id, first_name, last_name, email, type"
	SELECT_ID_BY_TG_ID = "select ac.id from accounts ac where ac.telegram_id = ($1) limit 1"
	SELECT_BY_TG_ID    = "select ac.* from accounts ac where ac.telegram_id = ($1)"
)

type SqlxAccountPort struct {
	Db          *sqlx.DB
	ProfileRepo intefraces.ProfileRepository
}

func (port *SqlxAccountPort) Register(entity authServerDomain.Account) (authServerDomain.Account, error) {
	var result authServerDomain.Account
	var resultErr error

	tx := port.Db.MustBegin()
	selErr := tx.Get(&result, SELECT_BY_TG_ID, int64(*entity.TelegramId))
	resultErr = repository.ProceedSelectErrorsWithCallback(selErr, tx)

	if result.TelegramId == nil && resultErr == nil {
		resultErr = tx.QueryRowx(INSERT_ACCOUNT, entity.FirstName, entity.LastName, entity.Email, entity.Type, entity.TelegramId).StructScan(&result)
		resultErr = repository.ProceedInsertErrorsWithCallback(resultErr, tx)
	}
	if resultErr == nil {
		resultErr = port.ProfileRepo.CreateProfile(result, "")
	}

	resultErr = repository.CommitAndProceedErrors(tx, resultErr)
	return result, resultErr
}

func (port *SqlxAccountPort) Save(entity authServerDomain.Account) (authServerDomain.Account, error) {
	var account authServerDomain.Account
	var errResp error
	insertErr := port.Db.QueryRowx(INSERT_ACCOUNT, entity.FirstName, entity.LastName, entity.Email, entity.Type, entity.TelegramId).StructScan(&account)
	if insertErr != nil {
		errResp = errors.Join(repository.InsertError, errors.New("Wrap error: "+insertErr.Error()))
	}
	return account, errResp
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
func (port *SqlxAccountPort) GetByTgId(tgId int64) (authServerDomain.Account, error) {
	var resp authServerDomain.Account
	var errResp error
	err := port.Db.Get(&resp, SELECT_BY_TG_ID, tgId)
	if err != nil {
		errResp = errors.Join(repository.SelectError, errors.New("Wrap error: "+err.Error()))
	}
	return resp, errResp
}

func (port *SqlxAccountPort) CloseConnection() {
	port.Db.Close()
}

func CreateSqlxAccountPort(db *sqlx.DB) *SqlxAccountPort {
	return &SqlxAccountPort{Db: db}
}
