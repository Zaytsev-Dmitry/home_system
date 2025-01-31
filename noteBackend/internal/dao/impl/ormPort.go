package noteDaoPorts

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	noteDomain "noteBackendApp/internal/domain"
)

type OrmNotePort struct {
	db *gorm.DB
}

func (port *OrmNotePort) Save(entity noteDomain.TelegramAccount) noteDomain.TelegramAccount {
	port.db.Clauses(clause.Returning{}).Create(&entity)
	return entity
}

func (port *OrmNotePort) DeleteNotesByAccountId(accountId int) {
	port.db.Where("account_id = ?", accountId).Delete(&noteDomain.TelegramAccount{})
}

func (port *OrmNotePort) GetNotesByAccountId(accountId int) []noteDomain.Note {
	var result []noteDomain.Note
	port.db.Where("account_id = ?", accountId).Find(&result)
	return result
}

func (port *OrmNotePort) ExistByName(name string) bool {
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

func (port *OrmNotePort) GetByName(name string) noteDomain.TelegramAccount {
	var result noteDomain.TelegramAccount
	port.db.Where("name = ?", name).Find(&result)
	return result
}

func CreateOrmNotePort(db *gorm.DB) *OrmNotePort {
	return &OrmNotePort{db: db}
}
