package connectdb

import (
// "database/sql"
// "github.com/jackc/pgx/v4"
// "github.com/penkong/data4life/gateway/config"
// "github.com/penkong/data4life/gateway/db/pgdb"
// "log"
)

// var Pdb *pgdb.Repo

// func Setup(conf *config.Config) {

// Open connection to database in this case Postgres13
// conn, err := sql.Open(conf.DBDriver, conf.DBSource)
// if err != nil {
// 	log.Fatal("db not connected!!!", err)
// }

// conn, err := pgx.Connect(context.Background(), conf.DBSource)
// if err != nil {
// 	fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
// 	os.Exit(1)
// }
// defer conn.Close(context.Background())

// err = conn.Ping()
// if err != nil {
// 	log.Fatal("ping error!!!", err)
// }

// Pdb = pgdb.NewRepo(conn)
// }
