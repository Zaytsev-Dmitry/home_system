package account

import (
	"authBackend/internal/app/domain"
	"authBackend/internal/app/ports/out/dao/queries"

	"authBackend/internal/app/ports/out/keycloak"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type AccountRepositorySqlx struct {
	db *sqlx.DB
}

func (port *AccountRepositorySqlx) CreateAccountAndProfile(entity keycloak.KeycloakEntity, username string, tgId int64) (domain.Account, error) {
	var result domain.Account
	var resultErr error

	//tx := port.db.MustBegin()
	//defer tx.Rollback()
	//
	//selectErr := tx.Get(&result, queries.SELECT_BY_TG_ID, tgId)
	//if selectErr == sql.ErrNoRows {
	//	//INSERT аккаунт
	//	errInsertAcc := tx.QueryRowx(queries.INSERT_ACCOUNT, entity.FirstName, entity.LastName, username, entity.Username, entity.Email, tgId, entity.ID, true).StructScan(&result)
	//	if errInsertAcc != nil {
	//		return domain.Account{}, repository.Fail(utilities.Error{
	//			Msg: "CreateAccountAndProfile.INSERT_ACCOUNT fail",
	//			Err: errInsertAcc,
	//		})
	//	}
	//	//создани базового профиля
	//	errInsertProf := tx.QueryRowx(profile.INSERT_PROFILE, result.ID, "USER", true, result.Username).Err()
	//	if errInsertProf != nil {
	//		return domain.Account{}, repository.Fail(utilities.Error{
	//			Msg: "CreateAccountAndProfile.INSERT_PROFILE fail",
	//			Err: errInsertAcc,
	//		})
	//	}
	//} else if selectErr != nil {
	//	return domain.Account{}, repository.Fail(utilities.Error{
	//		Msg: "CreateAccountAndProfile.SELECT_BY_TG_ID failed",
	//		Err: selectErr,
	//	})
	//}
	//if resultErr = tx.Commit(); resultErr != nil {
	//	resultErr = repository.Fail(utilities.Error{
	//		Msg: "CreateAccountAndProfile.Transaction commit fail",
	//		Err: resultErr,
	//	})
	//}
	return result, resultErr
}

func (port *AccountRepositorySqlx) GetIdByTgId(tgId int64) int64 {
	var resp int64
	err := port.db.Get(&resp, queries.SELECT_ID_BY_TG_ID, tgId)
	if err != nil {
		//TODO кинуть ошибку и потом создать базовый профиль
		fmt.Println(err)
	}
	return resp
}

func (port *AccountRepositorySqlx) GetByTgId(tgId int64) (domain.Account, error) {
	var resp domain.Account
	var errResp error
	//err := port.db.Get(&resp, queries.SELECT_BY_TG_ID, tgId)
	//if err != nil {
	//	if errors.Is(err, sql.ErrNoRows) {
	//		return domain.Account{}, repository.NoRows
	//	}
	//	errResp = errors.Join(repository.SelectError, errors.New("Wrap error: "+err.Error()))
	//}
	return resp, errResp
}

func NewAccountRepositorySqlx(db *sqlx.DB) *AccountRepositorySqlx {
	return &AccountRepositorySqlx{db: db}
}
