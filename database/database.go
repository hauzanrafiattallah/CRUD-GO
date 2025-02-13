package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func InitDatabase() *sql.DB {
	// Initialize database connection
	dsn := "root:Hauzanrafi09@tcp(localhost:3306)/CRUD_GO"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	// Check if the connection is successful
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
