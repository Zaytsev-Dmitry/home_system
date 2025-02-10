package domain

type UserAction struct {
	ID                uint64 `db:"id"`
	TelegramUserId    int64  `db:"telegram_user_id"`
	LastAction        string `db:"last_action"`
	LastRequirement   string `db:"last_requirement"`
	LastSentMessageId int    `db:"last_sent_message_id"`
}
