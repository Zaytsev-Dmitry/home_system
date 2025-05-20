package persistence

const (
	BOARD_INSERT                         = "insert into board (owner_id, name, currency) values($1, $2, $3) RETURNING *"
	BOARD_SELECT_ALL_BY_OWNER_ID         = "select b.* from board b where b.owner_id = ($1)"
	BOARD_SELECT_BY_ID                   = "select b.* from board b where b.id = ($1)"
	BOARD_PARTICIPANT_INSERT             = "insert into board_participant (board_id, participant_id) values($1, $2)"
	PARTICIPANT_SELECT_ID_BY_TG_USER_ID  = "select p.id from participant p where telegram_id = ($1)"
	PARTICIPANT_SELECT_IDS_BY_TG_USER_ID = "select p.id from participant p where telegram_id in (?)"
)
