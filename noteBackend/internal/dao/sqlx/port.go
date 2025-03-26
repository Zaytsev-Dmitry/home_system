package sqlx

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"noteBackendApp/internal/domain"
)

const (
	SELECT_BY_TELEGRAM_ID = "select * from notes where telegram_id=($1)"
	INSERT                = "insert into notes (account_id, telegram_id, name, link, description) values ($1, $2, $3, $4, $5) RETURNING *"
)

type SqlxAuthPort struct {
	Db *sqlx.DB
}

// TODO отловить ошибки
func (p *SqlxAuthPort) Save(entity domain.Note) domain.Note {
	var result domain.Note
	err := p.Db.QueryRowx(INSERT, entity.AccountId, entity.TelegramId, entity.Name, entity.Link, entity.Description).StructScan(&result)
	if err != nil {
		fmt.Println(err)
	}
	return result
}

func (p *SqlxAuthPort) DeleteNotesByTgId(tgId int64) {

}

func (p *SqlxAuthPort) GetNotesByTgId(tgId int64) []domain.Note {
	var resp []domain.Note
	err := p.Db.Select(&resp, SELECT_BY_TELEGRAM_ID, tgId)
	if err != nil {
		fmt.Println(err)
	}
	return resp
}

func (p *SqlxAuthPort) ExistByName(name string) bool {
	return false
}

func (p *SqlxAuthPort) GetByName(name string) domain.Note {
	return domain.Note{}
}

func (port *SqlxAuthPort) CloseConnection() {
	port.Db.Close()
}
