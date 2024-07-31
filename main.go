package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/nochzato/go-bank/api"
	db "github.com/nochzato/go-bank/db/sqlc"
	"github.com/nochzato/go-bank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalf("cannot read config file: %v", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalf("cannot connect to the db: %v", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatalf("cannot start server: %v", err)
	}
}
