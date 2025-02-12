package action

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"telegramCLient/internal/domain"
)

const (
	INSERT_ACTION   = "insert into user_action (telegram_user_id, need_user_input, last_sent_message_id,command_name, is_running) values($1, $2, $3, $4, $5) RETURNING id, telegram_user_id, need_user_input, last_sent_message_id, command_name, is_running"
	SELECT_BY_TG_ID = "select uc.* from user_action uc where uc.telegram_user_id = ($1)"
	UPDATE_BY_TG_ID = "update user_action set need_user_input = ($1), last_sent_message_id = ($2), command_name = ($3), is_running = ($4) where telegram_user_id = ($5)"
)

type SqlxActionPort struct {
	Db *sqlx.DB
}

func CreateSqlxActionPort(db *sqlx.DB) *SqlxActionPort {
	return &SqlxActionPort{Db: db}
}

// TODO отловаить ошибки
func (port *SqlxActionPort) SaveOrUpdate(telegramUserId int64, userInput bool, lastSentMessageId int, commandName string, isRunning bool) {
	var result domain.UserAction
	//var resultErr error

	tx := port.Db.MustBegin()
	defer tx.Rollback()

	err := tx.Get(&result, SELECT_BY_TG_ID, telegramUserId)
	if err != nil {
		fmt.Println(err)
	}
	if result.ID != 0 {
		err := tx.QueryRowx(UPDATE_BY_TG_ID, userInput, lastSentMessageId, commandName, isRunning, telegramUserId)
		if err.Err() != nil {
			fmt.Println(err)
		}
	} else {
		err := tx.QueryRowx(INSERT_ACTION, telegramUserId, userInput, lastSentMessageId, commandName, isRunning).StructScan(&result)
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
func (port *SqlxActionPort) Update(telegramUserId int64, needUserAction bool, lastSentMessageId int, isRunning bool) {
	port.Db.QueryRowx(UPDATE_BY_TG_ID, needUserAction, lastSentMessageId, telegramUserId, isRunning)
}

func (p *SqlxActionPort) CloseConnection() {
	p.Db.Close()
}
