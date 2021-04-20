package main

import (
	"html/template"
	"log"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

type Produto struct {
	Nome       string
	Preco      int64
	Descricao  string
	Quantidade int64
}

func main() {
	http.HandleFunc("/", handlePageWeb)
	log.Println("Servidor rodando...")
	http.ListenAndServe(":3333", nil)
}

func handlePageWeb(w http.ResponseWriter, r *http.Request) {
	p := []Produto{
		{"Base", 13, "Passar no rosto.", 10},
		{"Pó Banana", 13, "Passar no rosto.", 7},
		{"Baton", 3, "Passar no lábio.", 14},
		{"Mascara", 5, "Passar no rosto.", 4},
	}
	temp.ExecuteTemplate(w, "Index", p)
}
