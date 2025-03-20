package handlers

import (
	"go-reports/database"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func AddHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		content := strings.TrimSpace(r.FormValue("content"))
		_, err := database.DB.Exec("INSERT INTO reports (content) VALUES (?)", content)
		if err != nil {
			log.Fatal(err)
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		id := r.FormValue("id")
		_, err := database.DB.Exec("DELETE FROM reports WHERE id = ?", id)
		if err != nil {
			log.Fatal(err)
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		id := r.FormValue("id")
		content := strings.TrimSpace(r.FormValue("content"))
		_, err := database.DB.Exec("UPDATE reports SET content = ? WHERE id = ?", content, id)
		if err != nil {
			log.Fatal(err)
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func EditHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var report database.Report
	err := database.DB.QueryRow("SELECT id, content FROM reports WHERE id = ?", id).Scan(&report.ID, &report.Content)
	if err != nil {
		http.Error(w, "Report não encontrado", http.StatusNotFound)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/edit.html"))
	tmpl.Execute(w, report)
}
