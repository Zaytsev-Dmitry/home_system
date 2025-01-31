package authServerDomain

type Account struct {
	ID         uint64  `db:"id"`
	FirstName  *string `db:"first_name"`
	LastName   *string `db:"last_name"`
	Email      string  `db:"email"`
	Type       string  `db:"type"`
	TelegramId int     `db:"telegram_id"`
	IsActive   bool    `db:"is_active"`
}

type Profile struct {
	ID               uint64 `db:"id"`
	AccountId        uint64 `db:"account_id"`
	Role             string `db:"role"`
	TelegramUsername string `db:"telegram_username"`
}
