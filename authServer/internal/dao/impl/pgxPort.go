package authDaoPorts

import (
	authServerDomain "authServer/internal/domain"
	"github.com/jackc/pgx/v5"
)

type PgxAuthPort struct {
	Conn *pgx.Conn
}

func (port *PgxAuthPort) Save(entity authServerDomain.Account) authServerDomain.Account {
	return authServerDomain.Account{}
}

func CreatePgxAuthPort(conn *pgx.Conn) *PgxAuthPort {
	return &PgxAuthPort{Conn: conn}
}
