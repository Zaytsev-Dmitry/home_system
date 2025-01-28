package authDaoPorts

import (
	authServerDomain "authServer/internal/domain"
	"context"
	"github.com/jackc/pgx/v5"
)

type PgxAuthPort struct {
	Conn *pgx.Conn
}

func (port *PgxAuthPort) Save(entity authServerDomain.Account) authServerDomain.Account {
	return authServerDomain.Account{}
}

func (port *PgxAuthPort) CloseConnection() {
	port.Conn.Close(context.Background())
}

func CreatePgxAuthPort(conn *pgx.Conn) *PgxAuthPort {
	return &PgxAuthPort{Conn: conn}
}
