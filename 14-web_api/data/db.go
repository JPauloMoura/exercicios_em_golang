package data

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDb() *sql.DB {
	connectStr := "user=USER dbname=DB_NAME password=PASSWORD host=HOST sslmode=disable"

	connect, err := sql.Open("postgres", connectStr)
	if err != nil {
		log.Fatal(err)
	}
	return connect
}
