package main

import (
	"github.com/JPaulo-Moura/16-api-rest/database"

	"github.com/JPaulo-Moura/16-api-rest/server"
)

func main() {
	database.StartDB()
	s := server.NewServer()
	s.Run()
}
