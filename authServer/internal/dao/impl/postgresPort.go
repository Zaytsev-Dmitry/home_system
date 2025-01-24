package authDaoPorts

import (
	authServerDomain "authServer/internal/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PostgresAuthPort struct {
	db *gorm.DB
}

func CreatePostgresAuthPort(db *gorm.DB) *PostgresAuthPort {
	return &PostgresAuthPort{db: db}
}

func (port *PostgresAuthPort) Save(entity authServerDomain.Account) authServerDomain.Account {
	port.db.Clauses(clause.Returning{}).Create(&entity)
	return entity
}
