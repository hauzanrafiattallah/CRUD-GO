package controller

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

func NewUpdateUserController(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			id := r.URL.Query().Get("id")
			r.ParseForm()
			name := r.Form["name"][0]
			telephone := r.Form["telephone"][0]
			address := r.Form["address"][0]
			_, err := db.Exec("UPDATE users SET name=?, telephone=?, address=? WHERE id=?", name, telephone, address, id)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, "/user", http.StatusMovedPermanently)
			return
		} else if r.Method == "GET" {
			id := r.URL.Query().Get("id")

			row := db.QueryRow("SELECT id, name, telephone, address FROM users WHERE id = ?", id)

			if row.Err() != nil {
				w.Write([]byte(row.Err().Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			var user User

			err := row.Scan(
				&user.Id,
				&user.Name,
				&user.Telephone,
				&user.Address,
			)
			user.Id = id
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			fp := filepath.Join("views", "update.html")
			tmpl, err := template.ParseFiles(fp)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			data := make(map[string]any)
			data["user"] = user
			err = tmpl.Execute(w, data)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	}
}
