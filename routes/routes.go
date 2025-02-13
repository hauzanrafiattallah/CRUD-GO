package routes

import (
	"database/sql"
	"net/http"

	"github.com/hauzanrafiattallah/CRUD-GO/controller"
)

func MapRoutes(server *http.ServeMux, db *sql.DB) {
	server.HandleFunc("/", controller.NewHelloWorldController())
	server.HandleFunc("/user", controller.NewIndexUser())
	server.HandleFunc("/user/create", controller.NewCreateUserController(db))
}
