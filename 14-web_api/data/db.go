package data

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDb() *sql.DB {
	connectStr := "user=xxx dbname=xxx password=xxx host=xxx sslmode=xxx"

	connect, err := sql.Open("postgres", connectStr)
	if err != nil {
		log.Fatal(err)
	}
	return connect
}
