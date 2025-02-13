package routes

import (
	"net/http"

	"github.com/hauzanrafiattallah/CRUD-GO/controller"
)

func MapRoutes(server *http.ServeMux) {
	server.HandleFunc("/", controller.NewHelloWorldController())

}
