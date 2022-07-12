package main

import (
	app "github.com/penkong/fiber-sqlc/gateway"
	"log"
)

func main() {
	if err := app.Start(app.Conf.ServerAddress); err != nil {
		log.Fatalf("app.Start: %v\n", err)
	}
}
