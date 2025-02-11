package noteDaoPorts

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	noteDomain "noteBackendApp/internal/domain"
)

const SELECT_BY_TELEGRAM_ID = "select * from notes where telegram_id=($1)"

type SqlxAuthPort struct {
	Db *sqlx.DB
}

func (p *SqlxAuthPort) Save(entity noteDomain.Note) noteDomain.Note {
	return noteDomain.Note{}
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
