package participant

const (
	SELECT_ID_BY_TG_USER_ID = "select p.id from participant p where telegram_id = ($1)"
)
