package routes

import (
	"net/http"

	"github.com/14-web_api/controllers"
)

func Handler() {
	http.HandleFunc("/", controllers.HandlePageWeb)
}
