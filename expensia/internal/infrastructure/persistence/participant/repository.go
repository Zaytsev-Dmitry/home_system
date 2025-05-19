package participant

import (
	"database/sql"
	"errors"
	"github.com/Zaytsev-Dmitry/apikit/custom_errors"
	"github.com/jmoiron/sqlx"
)

type ParticipantRepositorySqlx struct {
	db *sqlx.DB
}

func NewParticipantRepositorySqlx(db *sqlx.DB) *ParticipantRepositorySqlx {
	return &ParticipantRepositorySqlx{db: db}
}

func (r *ParticipantRepositorySqlx) GetIdByTgUserId(userId int64) (int64, error) {
	var result int64
	err := r.db.Get(&result, SELECT_ID_BY_TG_USER_ID, userId)
	if errors.Is(err, sql.ErrNoRows) {
		return 0, custom_errors.RowNotFound
	}

	return result, err
}

func (r *ParticipantRepositorySqlx) GetIdByTgUserIdList(userIds []int64) ([]int64, error) {
	var result []int64
	query, args, err := sqlx.In(SELECT_IDS_BY_TG_USER_ID, userIds)
	if err != nil {
		return nil, err
	}
	query = r.db.Rebind(query)
	err = r.db.Select(&result, query, args...)
	if errors.Is(err, sql.ErrNoRows) || len(result) == 0 {
		return result, custom_errors.RowNotFound
	}
	return result, err
}
