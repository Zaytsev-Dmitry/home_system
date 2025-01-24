package noteDaoPorts

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	noteDomain "noteBackendApp/internal/domain"
)

type PostgresNotePort struct {
	db *gorm.DB
}

func (port *PostgresNotePort) Save(entity noteDomain.TelegramAccount) noteDomain.TelegramAccount {
	port.db.Clauses(clause.Returning{}).Create(&entity)
	return entity
}

func (port *PostgresNotePort) DeleteNotesByAccountId(accountId int) {
	port.db.Where("account_id = ?", accountId).Delete(&noteDomain.TelegramAccount{})
}

func (port *PostgresNotePort) GetNotesByAccountId(accountId int) []noteDomain.TelegramAccount {
	var result []noteDomain.TelegramAccount
	port.db.Where("account_id = ?", accountId).Find(&result)
	return result
}

func (port *PostgresNotePort) ExistByName(name string) bool {
	var exists bool
	err := port.db.Model(&noteDomain.TelegramAccount{}).
		Select("count(*) > 0").
		Where("name = ?", name).
		Find(&exists).
		Error
	if err != nil {
		fmt.Errorf(err.Error())
		return false
	}
	return exists
}

func (port *PostgresNotePort) GetByName(name string) noteDomain.TelegramAccount {
	var result noteDomain.TelegramAccount
	port.db.Where("name = ?", name).Find(&result)
	return result
}

func CreatePostgresNotePort(db *gorm.DB) *PostgresNotePort {
	return &PostgresNotePort{db: db}
}
