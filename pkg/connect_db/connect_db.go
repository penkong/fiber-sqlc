package connectdb

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/penkong/fiber-sqlc/config"
	"github.com/penkong/fiber-sqlc/db/pgdb"
	"log"
	"os"
)

var Pdb *pgdb.Repo

func Setup(conf *config.Config) {

	// Open connection to database in this case Postgres13
	conn, err := pgx.Connect(context.Background(), conf.DBSource)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	err = conn.Ping(context.Background())
	if err != nil {
		log.Fatal("ping error!!!", err)
	}

	Pdb = pgdb.NewRepo(conn)
}
