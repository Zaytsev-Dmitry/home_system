package domain

import "time"

type UserIdentityLink struct {
	ID             int64     `db:"id" json:"id"`
	Email          string    `db:"email" json:"email"`
	KeycloakUserID string    `db:"keycloak_user_id" json:"keycloakUserId"`
	TelegramUserID int64     `db:"telegram_user_id" json:"telegramUserId"`
	CreatedAt      time.Time `db:"created_at" json:"createdAt"`
}

type Profile struct {
	ID        uint64 `db:"id"`
	AccountId uint64 `db:"account_id"`
	Role      string `db:"role"`
	IsActive  bool   `db:"is_active"`
	Username  string `db:"username"`
}
