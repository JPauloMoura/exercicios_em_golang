package models

import (
	"log"

	"github.com/14-web_api/data"
)

type Product struct {
	Nome       string
	Preco      float64
	Descricao  string
	Quantidade int
}

func GetAllProducts() []Product {
	db := data.ConnectDb()
	defer db.Close()

	data, err := db.Query("select * from produtos")
	if err != nil {
		log.Fatal(err)
	}

	listProduct := []Product{}

	for data.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err := data.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			log.Fatal(err)
		}

		p := Product{nome, float64(preco), descricao, int(quantidade)}

		listProduct = append(listProduct, p)
	}
	return listProduct
}

func InsertProduct(p Product) {
	db := data.ConnectDb()
	defer db.Close()

	queryInsert, err := db.Prepare(`
		insert into produtos(nome, descricao, preco, quantidade)
		values($1, $2, $3, $4)
	`)

	if err != nil {
		panic(err.Error())
	}

	queryInsert.Exec(p.Nome, p.Descricao, p.Preco, p.Quantidade)
}
