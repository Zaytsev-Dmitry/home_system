package domain

import "time"

type Participant struct {
	Id             int64     `db:"id" json:"id"`
	KeycloakUserId string    `db:"keycloak_user_id" json:"keycloakUserId"` //unique not null
	TelegramId     int64     `db:"telegram_id" json:"telegramId"`          //unique not null
	DisplayName    string    `db:"display_name" json:"displayName,omitempty"`
	Username       string    `db:"username" json:"username,omitempty"`
	IsActive       bool      `db:"is_active" json:"isActive"`          //not null
	CreatedAtUtc   time.Time `db:"created_at_utc" json:"createdAtUtc"` //not null
}
