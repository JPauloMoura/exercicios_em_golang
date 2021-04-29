package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	_ "github.com/lib/pq"
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
	http.ListenAndServe(":3001", nil)
}

func handlePageWeb(w http.ResponseWriter, r *http.Request) {
	db := connectDb()
	defer db.Close()

	data, err := db.Query("select * from produtos")
	handlerErro(err)

	listProduct := []Produto{}

	for data.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err := data.Scan(&id, &nome, &descricao, &preco, &quantidade)
		handlerErro(err)

		p := Produto{nome, int64(preco), descricao, int64(quantidade)}
		listProduct = append(listProduct, p)
	}

	temp.ExecuteTemplate(w, "Index", listProduct)
}

func connectDb() *sql.DB {
	connStr := "user=DB_USER dbname=DB_NAME password=DB_PASSWORD host=DB_HOST sslmode=disable"
	connect, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	return connect
}

func handlerErro(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
