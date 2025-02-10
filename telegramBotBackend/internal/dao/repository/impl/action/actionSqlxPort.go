package action

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"telegramCLient/internal/domain"
)

const (
	INSERT_ACTION   = "insert into user_action (telegram_user_id, last_action, last_requirement, last_sent_message_id) values($1, $2, $3, $4) RETURNING id, telegram_user_id, last_action, last_requirement, last_sent_message_id"
	SELECT_BY_TG_ID = "select uc.* from user_action uc where uc.telegram_user_id = ($1)"
	UPDATE_BY_TG_ID = "update user_action set last_action = ($1), last_requirement = ($2), last_sent_message_id = ($3) where telegram_user_id = ($4)"
)

type SqlxActionPort struct {
	Db *sqlx.DB
}

func CreateSqlxActionPort(db *sqlx.DB) *SqlxActionPort {
	return &SqlxActionPort{Db: db}
}

// TODO отловаить ошибки
func (port *SqlxActionPort) Save(telegramUserId int64, lastAction string, lastRequirement string, lastSentMessageId int) {
	var result domain.UserAction
	//var resultErr error

	tx := port.Db.MustBegin()
	defer tx.Rollback()

	tx.Get(&result, SELECT_BY_TG_ID, telegramUserId)
	if result.ID != 0 {
		err := tx.QueryRowx(UPDATE_BY_TG_ID, lastAction, lastRequirement, lastSentMessageId, telegramUserId)
		if err.Err() != nil {
			fmt.Println(err)
		}
	} else {
		err := tx.QueryRowx(INSERT_ACTION, telegramUserId, lastAction, lastRequirement, lastSentMessageId).StructScan(&result)
		if err != nil {
			fmt.Println(err)
		}
	}
	tx.Commit()
}

// TODO отловаить ошибки
func (port *SqlxActionPort) GetByTgId(telegramUserId int64) domain.UserAction {
	var result domain.UserAction
	port.Db.Get(&result, SELECT_BY_TG_ID, telegramUserId)
	return result
}

// TODO отловаить ошибки
func (port *SqlxActionPort) Update(telegramUserId int64, lastAction string, lastRequirement string, lastSentMessageId int) {
	port.Db.QueryRowx(UPDATE_BY_TG_ID, lastAction, lastRequirement, lastSentMessageId, telegramUserId)
}

func (p *SqlxActionPort) CloseConnection() {
	p.Db.Close()
}
