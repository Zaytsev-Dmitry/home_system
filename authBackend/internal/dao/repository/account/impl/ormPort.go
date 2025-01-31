package impl

import (
	authServerDomain "authServer/internal/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type OrmAuthPort struct {
	Db *gorm.DB
}

func (port *OrmAuthPort) Save(entity authServerDomain.Account) authServerDomain.Account {
	port.Db.Clauses(clause.Returning{}).Create(&entity)
	return entity
}

func (port *OrmAuthPort) CloseConnection() {
	dbInstance, _ := port.Db.DB()
	dbInstance.Close()
}

func CreateOrmAuthPort(db *gorm.DB) *OrmAuthPort {
	return &OrmAuthPort{Db: db}
}
