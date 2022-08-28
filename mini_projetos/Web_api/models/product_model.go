package models

import (
	"log"
	"strconv"

	"github.com/14-web_api/data"
)

type Product struct {
	Id         int
	Nome       string
	Preco      float64
	Descricao  string
	Quantidade int
}

func GetAllProducts() []Product {
	db := data.ConnectDb()
	defer db.Close()

	query, err := db.Query("SELECT * FROM produtos ORDER BY nome ASC")
	if err != nil {
		log.Fatal(err)
	}

	listProduct := []Product{}

	for query.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err := query.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			log.Fatal(err)
		}

		p := Product{id, nome, float64(preco), descricao, int(quantidade)}

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

func DeleteProduct(id string) {
	db := data.ConnectDb()
	defer db.Close()

	query, err := db.Prepare(`
		DELETE FROM produtos WHERE id=$1
	`)
	if err != nil {
		panic("DeleteProduct: " + err.Error())
	}

	_, err = query.Exec(id)

	if err != nil {
		panic(err.Error())
	}
}

func GetProduct(id string) Product {
	db := data.ConnectDb()
	defer db.Close()

	intID, err := strconv.Atoi(string(id))
	if err != nil {
		panic("EditProduct:" + err.Error())
	}

	query, err := db.Query(`SELECT * FROM produtos WHERE id=$1`, intID)

	if err != nil {
		panic("EditProduct: " + err.Error())
	}

	p := Product{}

	for query.Next() {
		err := query.Scan(&p.Id, &p.Nome, &p.Descricao, &p.Preco, &p.Quantidade)
		if err != nil {
			log.Fatal(err)
		}
	}

	return p
}

func UpdateProduct(p Product) {
	db := data.ConnectDb()
	defer db.Close()

	// product := GetProduct(string(p.Id))

	query, err := db.Prepare(`
		UPDATE produtos SET 
		nome=$1, preco=$2, descricao=$3, quantidade=$4
		WHERE id=$5
	`)
	if err != nil {
		log.Println("UpdateProduct" + err.Error())
	}

	_, err = query.Exec(p.Nome, p.Preco, p.Descricao, p.Quantidade, p.Id)

	if err != nil {
		log.Println("UpdateProduct", err.Error())
	}

}
