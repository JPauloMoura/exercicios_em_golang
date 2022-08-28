package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/14-web_api/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	listProduct := models.GetAllProducts()
	temp.ExecuteTemplate(w, "Index", listProduct)
}

func HandlerForm(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Form", nil)
}

func InsertProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		precoFloat, err := strconv.ParseFloat(r.FormValue("preco"), 64)
		if err != nil {
			log.Println("InsertProduct: conversão de preço", err)
		}

		quantidadeInt, err := strconv.Atoi(r.FormValue("quantidade"))
		if err != nil {
			log.Println("InsertProduct: conversão de quantidade", err)
		}

		p := models.Product{
			Nome:       r.FormValue("nome"),
			Descricao:  r.FormValue("descricao"),
			Preco:      precoFloat,
			Quantidade: quantidadeInt,
		}

		models.InsertProduct(p)
	}

	http.Redirect(w, r, "/", 301)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.URL.Query().Get("id")

	models.DeleteProduct(productID)

	http.Redirect(w, r, "/", 301)

}

func EditProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.URL.Query().Get("id")
	product := models.GetProduct(productID)
	temp.ExecuteTemplate(w, "Edit", product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		idInt, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			log.Println("UpdateProduct id:", err.Error())
		}

		precoFloat, err := strconv.ParseFloat(r.FormValue("preco"), 64)
		if err != nil {
			log.Println("UpdateProduct preco:", err.Error())
		}

		quantidadeInt, err := strconv.Atoi(r.FormValue("quantidade"))
		if err != nil {
			log.Println("UpdateProduct quantidade:", err.Error())
		}

		p := models.Product{
			Id:         idInt,
			Nome:       r.FormValue("nome"),
			Preco:      precoFloat,
			Descricao:  r.FormValue("descricao"),
			Quantidade: quantidadeInt,
		}

		models.UpdateProduct(p)
	}
	http.Redirect(w, r, "/", 301)
}
