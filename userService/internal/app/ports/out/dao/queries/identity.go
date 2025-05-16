package queries

const (
	INSERT_IDENTITY_LINK = "insert into user_identity_links (keycloak_user_id, telegram_user_id, email) values($1, $2, $3) RETURNING *"
)
