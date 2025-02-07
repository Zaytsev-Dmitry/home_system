package account

import (
	"errors"
	"github.com/jmoiron/sqlx"
)

func (port *SqlxAccountPort) proceedSelectErrorsWithCallback(resultErr error, tx *sqlx.Tx) error {
	if resultErr != nil {
		resultErr = SelectError
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			resultErr = errors.Join(SelectError, errors.New("Wrap error: "+rollbackErr.Error()))
		}
	}
	return resultErr
}

func (port *SqlxAccountPort) proceedInsertErrorsWithCallback(resultErr error, tx *sqlx.Tx) error {
	if resultErr != nil {
		resultErr = InsertError
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			resultErr = errors.Join(InsertError, errors.New("Wrap error: "+rollbackErr.Error()))
		}
	}
	return resultErr
}

func (port *SqlxAccountPort) commitAndProceedErrors(tx *sqlx.Tx, resultErr error) error {
	commitErr := tx.Commit()
	if commitErr != nil {
		resultErr = errors.Join(CommitError, errors.New("Wrap error: "+commitErr.Error()))
	}
	return resultErr
}
