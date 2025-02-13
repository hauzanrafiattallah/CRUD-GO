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
			nik := r.Form["nik"][0]
			address := r.Form["address"][0]
			_, err := db.Exec("UPDATE user SET name=?, nik=?, address=? WHERE id=?", name, nik, address, id)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, "/user", http.StatusMovedPermanently)
			return
		} else if r.Method == "GET" {
			id := r.URL.Query().Get("id")

			row := db.QueryRow("SELECT id, name, nik, address FROM user WHERE id = ?", id)

			if row.Err() != nil {
				w.Write([]byte(row.Err().Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			var user User

			err := row.Scan(
				&user.ID,
				&user.Name,
				&user.Nik,
				&user.Address,
			)
			user.ID = id
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
