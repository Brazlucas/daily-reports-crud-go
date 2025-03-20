package handlers

import (
	"go-reports/database"
	"html/template"
	"log"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	reports := []database.Report{}
	rows, err := database.DB.Query("SELECT id, content, timestamp FROM reports ORDER BY timestamp DESC")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var report database.Report
		rows.Scan(&report.ID, &report.Content, &report.Timestamp)
		reports = append(reports, report)
	}

	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, reports)
}
