package dao_helper

import (
	customErr "authBackend/pkg/errors"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
)

const (
	Get       = "Get"
	Exec      = "Exec"
	QueryRowx = "QueryRowx"
)

func ExecuteQuery[T any](queryType string, db *sqlx.DB, query, action string, args ...interface{}) (*T, *customErr.CustomError) {
	var err error
	var result T

	switch queryType {
	case Get:
		err = db.Get(&result, query, args...)
	case QueryRowx:
		err = db.QueryRowx(query, args...).StructScan(&result)
	}

	if err != nil {
		return nil, customErr.New(action, err)
	}
	return &result, nil
}

func ExecuteQueryWithOutEntityResponse(db *sqlx.DB, query, action string, args ...interface{}) *customErr.CustomError {
	var err error
	_, err = db.Exec(query, args...)
	if err != nil {
		return customErr.New(action, err)
	}
	return nil
}

func ExecuteQuerySlice[T any](db *sqlx.DB, query, action string, args ...interface{}) ([]*T, *customErr.CustomError) {
	var results []T
	err := db.Select(&results, query, args...)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []*T{}, nil // Возвращаем пустой срез указателей
		}
		return nil, customErr.New(action, err)
	}

	// Преобразуем []T → []*T
	ptrResults := make([]*T, len(results))
	for i := range results {
		ptrResults[i] = &results[i]
	}

	return ptrResults, nil
}
