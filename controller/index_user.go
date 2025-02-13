package controller

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

type User struct {
	Id        string
	Name      string
	Telephone string
	Address   string
}

func NewIndexUser(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, name, telephone, address FROM users")
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer rows.Close()
		var users []User
		for rows.Next() {
			var user User
			err = rows.Scan(&user.Id, &user.Name, &user.Telephone, &user.Address)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			users = append(users, user)
		}
		fp := filepath.Join("views", "index.html")
		tmpl, err := template.ParseFiles(fp)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		data := make(map[string]any)
		data["users"] = users
		err = tmpl.Execute(w, data)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
