package main

import (
	"log"
	"net/http"

	"github.com/14-web_api/routes"
)

func main() {
	routes.Handler()
	log.Println("Servidor rodando na porta 3002...")
	http.ListenAndServe(":3002", nil)
}
