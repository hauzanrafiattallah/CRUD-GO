package routes

import (
	"database/sql"
	"net/http"

	"github.com/hauzanrafiattallah/CRUD-GO/controller"
)

func MapRoutes(server *http.ServeMux, db *sql.DB) {
	server.HandleFunc("/", controller.NewHelloWorldController())
	server.HandleFunc("/user", controller.NewIndexUser(db))
	server.HandleFunc("/user/create", controller.NewCreateUserController(db))
	server.HandleFunc("/user/update", controller.NewUpdateUserController(db))
	server.HandleFunc("/user/delete", controller.NewDeleteUserController(db))
}
