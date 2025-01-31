package noteDaoPorts

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	noteDomain "noteBackendApp/internal/domain"
)

const SELECT_BY_NAME = "select * from notes where account_id=($1)"

type SqlxAuthPort struct {
	Db *sqlx.DB
}

func (p *SqlxAuthPort) Save(entity noteDomain.TelegramAccount) noteDomain.TelegramAccount {
	return noteDomain.TelegramAccount{}
}
func (p *SqlxAuthPort) DeleteNotesByAccountId(accountId int) {

}

func (p *SqlxAuthPort) GetNotesByAccountId(accountId int) []noteDomain.Note {
	var resp []noteDomain.Note
	err := p.Db.Select(&resp, SELECT_BY_NAME, accountId)
	if err != nil {
		fmt.Println(err)
	}
	return resp
}

func (p *SqlxAuthPort) ExistByName(name string) bool {
	return false
}

func (p *SqlxAuthPort) GetByName(name string) noteDomain.TelegramAccount {
	return noteDomain.TelegramAccount{}
}

func (port *SqlxAuthPort) CloseConnection() {
	port.Db.Close()
}

func CreateSqlxAuthPort(db *sqlx.DB) *SqlxAuthPort {
	return &SqlxAuthPort{Db: db}
}
