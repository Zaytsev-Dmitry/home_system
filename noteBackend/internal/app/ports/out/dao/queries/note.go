package queries

const (
	ExistByName     = "select * from notes where name = $1"
	Insert          = "INSERT INTO notes (name, link, description, telegram_id) VALUES ($1, $2, $3, $4) RETURNING *"
	GetAllByTgId    = "select * from notes where telegram_id = $1"
	DeleteAllByTgId = "delete from notes where telegram_id = $1"
)
