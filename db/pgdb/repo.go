package pgdb

import "github.com/jackc/pgx/v4"

type Repo struct {
	*Queries
	DB *pgx.Conn
}

func NewRepo(db *pgx.Conn) *Repo {
	return &Repo{
		DB:      db,
		Queries: New(db),
	}
}
