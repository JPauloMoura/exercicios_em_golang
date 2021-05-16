package routes

import (
	"net/http"

	"github.com/14-web_api/controllers"
)

func Handler() {
	http.HandleFunc("/", controllers.HandleIndex)
	http.HandleFunc("/novo-produto", controllers.HandlerForm)
	http.HandleFunc("/insert", controllers.InsertProduct)
	http.HandleFunc("/delete", controllers.DeleteProduct)
	http.HandleFunc("/editar", controllers.EditProduct)
	http.HandleFunc("/update", controllers.UpdateProduct)
}
