package main

import (
	"github.com/JPaulo-Moura/16-api-rest/server"
)

func main() {
	s := server.NewServer()
	s.Run()
}
