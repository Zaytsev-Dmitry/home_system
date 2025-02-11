package account

import (
	"authServer/external/keycloak"
	"authServer/internal/dao/repository"
	"authServer/internal/dao/repository/impl/profile"
	"authServer/internal/dao/repository/intefraces"
	authServerDomain "authServer/internal/domain"
	"authServer/pkg/utilities"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	INSERT_ACCOUNT     = "insert into accounts (first_name, last_name, username, telegram_user_name, email, telegram_id, keycloak_id, is_active) values($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id, first_name, last_name, username,telegram_user_name, email, telegram_id, keycloak_id, is_active"
	SELECT_ID_BY_TG_ID = "select ac.id from accounts ac where ac.telegram_id = ($1) limit 1"
	SELECT_BY_TG_ID    = "select ac.* from accounts ac where ac.telegram_id = ($1)"
)

type SqlxAccountPort struct {
	Db          *sqlx.DB
	ProfileRepo intefraces.ProfileRepository
}

func (port *SqlxAccountPort) CreateAccountAndProfile(entity keycloak.KeycloakEntity, username string, tgId int64) (authServerDomain.Account, error) {
	var result authServerDomain.Account
	var resultErr error

	tx := port.Db.MustBegin()
	defer tx.Rollback()

	selectErr := tx.Get(&result, SELECT_BY_TG_ID, tgId)
	if selectErr == sql.ErrNoRows {
		//INSERT аккаунт
		errInsertAcc := tx.QueryRowx(INSERT_ACCOUNT, entity.FirstName, entity.LastName, username, entity.Username, entity.Email, tgId, entity.ID, true).StructScan(&result)
		if errInsertAcc != nil {
			return authServerDomain.Account{}, repository.Fail(utilities.Error{
				Msg: "CreateAccountAndProfile.INSERT_ACCOUNT fail",
				Err: errInsertAcc,
			})
		}
		//создани базового профиля
		errInsertProf := tx.QueryRowx(profile.INSERT_PROFILE, result.ID, "USER", true, result.Username).Err()
		if errInsertProf != nil {
			return authServerDomain.Account{}, repository.Fail(utilities.Error{
				Msg: "CreateAccountAndProfile.INSERT_PROFILE fail",
				Err: errInsertAcc,
			})
		}
	} else if selectErr != nil {
		return authServerDomain.Account{}, repository.Fail(utilities.Error{
			Msg: "CreateAccountAndProfile.SELECT_BY_TG_ID failed",
			Err: selectErr,
		})
	}
	if resultErr = tx.Commit(); resultErr != nil {
		resultErr = repository.Fail(utilities.Error{
			Msg: "CreateAccountAndProfile.Transaction commit fail",
			Err: resultErr,
		})
	}
	return result, resultErr
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

func CreateSqlxAccountPort(db *sqlx.DB, prof *intefraces.ProfileRepository) *SqlxAccountPort {
	return &SqlxAccountPort{Db: db, ProfileRepo: *prof}
}
