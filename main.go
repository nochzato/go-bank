package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/nochzato/go-bank/api"
	db "github.com/nochzato/go-bank/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://postgres:postgres@localhost:5432/go_bank?sslmode=disable"
	serverAddress = "localhost:8000"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatalf("cannot connect to the db: %v", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatalf("cannot start server: %v", err)
	}
}
