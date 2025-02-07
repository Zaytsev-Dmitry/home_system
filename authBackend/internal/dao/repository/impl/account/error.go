package account

import (
	"errors"
	"github.com/jmoiron/sqlx"
)

func (port *SqlxAccountPort) proceedSelectErrorsWithCallback(err error, tx *sqlx.Tx) error {
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

func (port *SqlxAccountPort) proceedInsertErrorsWithCallback(err error, tx *sqlx.Tx) error {
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

func (port *SqlxAccountPort) commitAndProceedErrors(tx *sqlx.Tx, resultErr error) error {
	if resultErr == nil {
		commitErr := tx.Commit()
		if commitErr != nil {
			resultErr = errors.Join(CommitError, errors.New("Wrap error: "+commitErr.Error()))
		}
	}
	return resultErr
}
