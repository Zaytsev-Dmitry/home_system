package noteDaoPorts

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	noteDomain "noteBackendApp/internal/domain"
)

const (
	SELECT_BY_TELEGRAM_ID = "select * from notes where telegram_id=($1)"
	INSERT                = "insert into notes (account_id, telegram_id, name, link, description) values ($1, $2, $3, $4, $5) RETURNING *"
)

type SqlxAuthPort struct {
	Db *sqlx.DB
}

// TODO отловить ошибки
func (p *SqlxAuthPort) Save(entity noteDomain.Note) noteDomain.Note {
	var result noteDomain.Note
	err := p.Db.QueryRowx(INSERT, entity.AccountId, entity.TelegramId, entity.Name, entity.Link, entity.Description).StructScan(&result)
	if err != nil {
		fmt.Println(err)
	}
	return result
}

func (p *SqlxAuthPort) DeleteNotesByTgId(tgId int64) {

}

func (p *SqlxAuthPort) GetNotesByTgId(tgId int64) []noteDomain.Note {
	var resp []noteDomain.Note
	err := p.Db.Select(&resp, SELECT_BY_TELEGRAM_ID, tgId)
	if err != nil {
		fmt.Println(err)
	}
	return resp
}

func (p *SqlxAuthPort) ExistByName(name string) bool {
	return false
}

func (p *SqlxAuthPort) GetByName(name string) noteDomain.Note {
	return noteDomain.Note{}
}

func (port *SqlxAuthPort) CloseConnection() {
	port.Db.Close()
}

func CreateSqlxAuthPort(db *sqlx.DB) *SqlxAuthPort {
	return &SqlxAuthPort{Db: db}
}
