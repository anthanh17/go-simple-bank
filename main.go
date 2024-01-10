package main

import (
	"database/sql"
	"log"

	"github.com/anthanh17/simplebank/api"
	db "github.com/anthanh17/simplebank/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:abc123@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:9000"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}