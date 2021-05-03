package controllers

import (
	"html/template"
	"net/http"

	"github.com/14-web_api/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func HandlePageWeb(w http.ResponseWriter, r *http.Request) {
	listProduct := models.GetAllProducts()
	temp.ExecuteTemplate(w, "Index", listProduct)
}
