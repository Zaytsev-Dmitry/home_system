package domain

type UserAction struct {
	ID                uint64 `db:"id"`
	TelegramUserId    int64  `db:"telegram_user_id"`
	CommandState      string `db:"command_state"`
	NeedUserAction    bool   `db:"need_user_action"`
	LastSentMessageId int    `db:"last_sent_message_id"`
	CommandName       string `db:"command_name"`
}
