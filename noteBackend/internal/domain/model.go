package noteDomain

type Note struct {
	ID          uint   `db:"id"`
	AccountId   int    `db:"account_id"`
	TelegramId  int64  `db:"telegram_id"`
	Name        string `db:"name"`
	Link        string `db:"link"`
	Description string `db:"description"`
}
