package main

import (
	"net/http"

	"github.com/hauzanrafiattallah/CRUD-GO/database"
	"github.com/hauzanrafiattallah/CRUD-GO/routes"
)

func main() {
	db := database.InitDatabase()

	server := http.NewServeMux()

	routes.MapRoutes(server, db)

	http.ListenAndServe(":8080", server)
}
