package persistence

import (
	apikitErr "github.com/Zaytsev-Dmitry/apikit/custom_errors"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"log"
)

type BoardParticipantRepositorySqlx struct {
	db *sqlx.DB
}

func NewBoardParticipantRepositorySqlx(db *sqlx.DB) *BoardParticipantRepositorySqlx {
	return &BoardParticipantRepositorySqlx{db: db}
}

func (bp BoardParticipantRepositorySqlx) AddParticipantsToBoard(boardId int64, participantIds []int64) error {
	tx := bp.db.MustBegin()
	defer tx.Rollback()

	for _, id := range participantIds {
		_, err := tx.Exec(BOARD_PARTICIPANT_INSERT, boardId, id)
		if err != nil {
			if dbErr, ok := err.(*pq.Error); ok && dbErr.Code == "23505" {
				log.Printf("BoardRepositorySqlx.Save conflict: %s (Detail: %s)", dbErr.Code, dbErr.Detail)
				return apikitErr.ConflictError
			} else {
				return err
			}
		}
	}
	tx.Commit()
	return nil
}
