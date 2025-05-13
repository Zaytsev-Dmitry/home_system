package queries

const (
	INSERT_ACCOUNT     = "insert into accounts (first_name, last_name, username, telegram_user_name, email, telegram_id, keycloak_id, is_active) values($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id, first_name, last_name, username,telegram_user_name, email, telegram_id, keycloak_id, is_active"
	SELECT_ID_BY_TG_ID = "select ac.id from accounts ac where ac.telegram_id = ($1) limit 1"
	SELECT_BY_TG_ID    = "select ac.* from accounts ac where ac.telegram_id = ($1)"
)
