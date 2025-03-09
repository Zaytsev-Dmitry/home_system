package domain

type UserAction struct {
	ID                uint64 `db:"id"`
	TelegramUserId    int64  `db:"telegram_user_id"`
	NeedUserInput     bool   `db:"need_user_input"`
	LastSentMessageId int    `db:"last_sent_message_id"`
	CommandName       string `db:"command_name"`
	IsRunning         bool   `db:"is_running"`
}
