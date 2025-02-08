package repository

import (
	"authServer/pkg/utilities"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

var (
	SelectError = errors.New("SqlxAccountPort.Select error.")
	InsertError = errors.New("SqlxAccountPort.Insert error.")
	CommitError = errors.New("SqlxAccountPort.Tx commit error.")
)

func RollbackTx(inErr error, tx *sqlx.Tx) {
	if inErr != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			inErr = errors.Join(inErr, errors.New("Wrap error: "+rollbackErr.Error()))
		}
	}
}

func ProceedErrorWithRollback(err error, tx *sqlx.Tx) error {
	RollbackTx(err, tx)
	return err
}

func CommitOrRollbackTx(err []error, tx *sqlx.Tx) error {
	var res error
	if len(err) > 0 {
		for i, value := range err {
			if value != nil {
				if res == nil {
					res = err[i]
				} else {
					res = errors.Join(res, errors.New("Wrap error: "+value.Error()))
				}
			}
		}
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			res = errors.Join(res, errors.New("Wrap error: "+rollbackErr.Error()))
		}
	} else {
		tx.Commit()
	}
	return res
}

func CommitAndProceedErrors(tx *sqlx.Tx, resultErr error) error {
	if resultErr == nil {
		commitErr := tx.Commit()
		if commitErr != nil {
			resultErr = errors.Join(CommitError, errors.New("Wrap error: "+commitErr.Error()))
		}
	}
	return resultErr
}

func Fail(error utilities.Error) error {
	logger := utilities.GetLogger()
	text := error.Msg + "." + " Wrapped: " + error.Err.Error()
	logger.Error(text)
	return fmt.Errorf(text)
}
