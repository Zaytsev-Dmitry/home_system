package board

import (
	"database/sql"
	"errors"
	"expensia/internal/app/domain"
	"expensia/internal/app/usecases"
	apikit "github.com/Zaytsev-Dmitry/apikit/custom_errors"
	"github.com/Zaytsev-Dmitry/dbkit"
	"github.com/Zaytsev-Dmitry/dbkit/custom_error"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"log"
	"strconv"
)

type BoardRepositorySqlx struct {
	db *sqlx.DB
}

func NewBoardRepositorySqlx(db *sqlx.DB) *BoardRepositorySqlx {
	return &BoardRepositorySqlx{db: db}
}

func (b BoardRepositorySqlx) GetAllByTgUserId(ownerId int64) ([]*domain.Board, *custom_error.CustomError) {
	return dbkit.ExecuteQuerySlice[domain.Board](
		b.db,
		SELECT_ALL_BY_OWNER_ID,
		"Получение все досок по telegram user id",
		ownerId,
	)
}

func (b BoardRepositorySqlx) SaveAndFlush(req usecases.CreateBoardInput) (*domain.Board, error) {
	query, err := dbkit.ExecuteQuery[domain.Board](
		true,
		"QueryRowx",
		b.db,
		INSERT_BOARD,
		"Создание доски для пользователя с telegram id == "+strconv.FormatInt(req.TgUserId, 10)+" и именем: "+req.Name,
		req.OwnerId, req.Name, req.Currency,
	)
	if err != nil {
		if dbErr, ok := err.WrapErr.(*pq.Error); ok && dbErr.Code == "23505" {
			log.Printf("409 CONFLICT: %s: (Detail: %s)", err.Action, dbErr.Detail)
			return nil, apikit.ConflictError
		} else {
			return nil, err.WrapErr
		}
	}
	return query, nil
}

func (b BoardRepositorySqlx) GetById(boardId int64) (*domain.Board, error) {
	result, err := dbkit.ExecuteQuery[domain.Board](
		false,
		"Get",
		b.db,
		SELECT_BY_ID,
		"Получение доски по id == "+strconv.FormatInt(boardId, 10),
		boardId,
	)

	if err != nil && errors.Is(err.WrapErr, sql.ErrNoRows) {
		return nil, apikit.RowNotFound
	}

	return result, nil
}
