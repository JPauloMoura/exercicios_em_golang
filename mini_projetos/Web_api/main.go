package main

import (
	"log"
	"net/http"

	"github.com/14-web_api/routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	routes.Handler()
	log.Println("Servidor rodando na porta 3002...")
	http.ListenAndServe(":3002", nil)
}
