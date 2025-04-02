package note

import (
	"github.com/jmoiron/sqlx"
	"noteBackendApp/internal/app/domain"
	"noteBackendApp/internal/app/ports/out/dao/queries"
	"noteBackendApp/pkg/dao_helper"
	"noteBackendApp/pkg/errors"
	"strconv"
)

type NoteRepositorySqlx struct {
	db *sqlx.DB
}

func (n NoteRepositorySqlx) Save(entity domain.Note) (*domain.Note, *errors.CustomError) {
	return dao_helper.ExecuteQuery[domain.Note](
		dao_helper.QueryRowx,
		n.db,
		queries.Insert,
		"сохранение записки",
		entity.Name, entity.Link, entity.Description, entity.TelegramId,
	)
}

func (n NoteRepositorySqlx) DeleteNotesByTgId(tgId int64) *errors.CustomError {
	return dao_helper.ExecuteQueryWithOutEntityResponse(
		n.db,
		queries.DeleteAllByTgId,
		"получение всех записок по telegram_id: "+strconv.FormatInt(tgId, 10),
		tgId,
	)
}

func (n NoteRepositorySqlx) GetNotesByTgId(tgId int64) ([]*domain.Note, *errors.CustomError) {
	return dao_helper.ExecuteQuerySlice[domain.Note](
		n.db,
		queries.GetAllByTgId,
		"получение всех записок по telegram_id: "+strconv.FormatInt(tgId, 10),
		tgId,
	)
}

func (n NoteRepositorySqlx) ExistByName(name string) (*domain.Note, *errors.CustomError) {
	return dao_helper.ExecuteQuery[domain.Note](
		dao_helper.Get,
		n.db,
		queries.ExistByName,
		"проверка на существование записки по имени: "+name,
		name,
	)
}

func NewNoteSqlx(db *sqlx.DB) *NoteRepositorySqlx {
	return &NoteRepositorySqlx{db: db}
}
