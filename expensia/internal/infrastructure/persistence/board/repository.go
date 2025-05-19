package board

import (
	"expensia/internal/app/domain"
	"expensia/internal/app/ports/out/dao/repository"
	apikitErr "github.com/Zaytsev-Dmitry/apikit/custom_errors"
	"github.com/Zaytsev-Dmitry/dbkit"
	"github.com/Zaytsev-Dmitry/dbkit/custom_error"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"log"
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

func (b BoardRepositorySqlx) SaveAndFlush(req repository.CreateBoardUCaseIn) (*domain.Board, error) {
	var result domain.Board

	tx := b.db.MustBegin()
	defer tx.Rollback()

	err := tx.QueryRowx(INSERT_BOARD, req.OwnerId, req.Name, req.Currency).StructScan(&result)
	if err != nil {
		if dbErr, ok := err.(*pq.Error); ok && dbErr.Code == "23505" {
			log.Printf("BoardRepositorySqlx.Save conflict: %s (Detail: %s)", dbErr.Code, dbErr.Detail)
			return nil, apikitErr.ConflictError
		} else {
			return nil, err
		}
	} else {
		tx.Commit()
		return &result, err
	}
}

func (b BoardRepositorySqlx) GetById(boardId int64) (*domain.Board, error) {
	var result domain.Board
	err := b.db.Get(&result, SELECT_BY_ID, boardId)
	return &result, err
}
