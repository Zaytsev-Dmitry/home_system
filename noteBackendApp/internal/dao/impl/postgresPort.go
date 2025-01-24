package noteDaoPorts

import (
	"gorm.io/gorm"
	noteDomain "noteBackendApp/internal/domain"
)

type PostgresNotePort struct {
	db *gorm.DB
}

func (port *PostgresNotePort) Save(entity noteDomain.TelegramAccount) noteDomain.TelegramAccount {
	return noteDomain.TelegramAccount{}
}

func CreatePostgresAuthPort(db *gorm.DB) *PostgresNotePort {
	return &PostgresNotePort{db: db}
}
