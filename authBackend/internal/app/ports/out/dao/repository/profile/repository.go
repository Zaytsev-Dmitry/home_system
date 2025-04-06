package profile

import (
	"authBackend/internal/app/domain"
	"authBackend/internal/app/ports/out/dao/queries"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type ProfileRepositorySqlx struct {
	Db *sqlx.DB
}

func (p *ProfileRepositorySqlx) CreateProfile(account domain.Account) error {
	//var result domain.Profile
	//var resultErr error
	//
	//tx := p.Db.MustBegin()
	//if p.GetProfileByAccountId(account.ID).ID == 0 {
	//	insertErr := tx.QueryRowx(queries.INSERT_PROFILE, account.ID, "USER", true, account.Username).StructScan(&result)
	//	resultErr = repository.ProceedErrorWithRollback(insertErr, tx)
	//}
	//resultErr = repository.CommitAndProceedErrors(tx, resultErr)
	//return resultErr
	return nil
}

func (p *ProfileRepositorySqlx) GetProfileByAccountId(accId int64) domain.Profile {
	var resp domain.Profile
	err := p.Db.Get(&resp, queries.SELECT_BY_ACC_ID, accId)
	if err != nil {
		//TODO кинуть ошибку и потом создать базовый профиль
		fmt.Println(err)
	}
	return resp
}

func NewProfileRepositorySqlx(db *sqlx.DB) *ProfileRepositorySqlx {
	return &ProfileRepositorySqlx{Db: db}
}
