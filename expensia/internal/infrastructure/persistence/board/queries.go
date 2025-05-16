package board

const (
	INSERT_BOARD           = "insert into board (owner_id, name, currency) values($1, $2, $3) RETURNING *"
	SELECT_ALL_BY_OWNER_ID = "select b.* from board b where b.owner_id = ($1)"
)
