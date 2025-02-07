package repository

import (
	"errors"
	"github.com/jmoiron/sqlx"
)

var (
	SelectError = errors.New("SqlxAccountPort.Select error.")
	InsertError = errors.New("SqlxAccountPort.Insert error.")
	CommitError = errors.New("SqlxAccountPort.Tx commit error.")
)

func ProceedSelectErrorsWithCallback(err error, tx *sqlx.Tx) error {
	var resultErr error
	if err != nil {
		resultErr = errors.Join(SelectError, errors.New("Wrap error: "+err.Error()))
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			resultErr = errors.Join(SelectError, errors.New("Wrap error_1: "+err.Error()), errors.New("Wrap error_2: "+rollbackErr.Error()))
		}
	}
	return resultErr
}

func ProceedInsertErrorsWithCallback(err error, tx *sqlx.Tx) error {
	var resultErr error
	if err != nil {
		resultErr = errors.Join(InsertError, errors.New("Wrap error: "+err.Error()))
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			resultErr = errors.Join(InsertError, errors.New("Wrap error_1: "+err.Error()), errors.New("Wrap error_2: "+rollbackErr.Error()))
		}
	}
	return resultErr
}

func CommitAndProceedErrors(tx *sqlx.Tx, resultErr error) error {
	if resultErr == nil {
		commitErr := tx.Commit()
		if commitErr != nil {
			resultErr = errors.Join(CommitError, errors.New("Wrap error: "+commitErr.Error()))
		}
	} else {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			resultErr = errors.Join(InsertError, errors.New("Wrap error_1: "+resultErr.Error()), errors.New("Wrap error_2: "+rollbackErr.Error()))
		}
	}
	return resultErr
}
