package main

import (
	"fmt"
	"go-reports/database"
	"go-reports/handlers"
	"go-reports/utils"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	database.InitDB()
	defer database.DB.Close()

	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/add", handlers.AddHandler)
	mux.HandleFunc("/delete", handlers.DeleteHandler)
	mux.HandleFunc("/update", handlers.UpdateHandler)
	mux.HandleFunc("/edit", handlers.EditHandler)

	port := ":7777"
	fmt.Println("Servidor rodando em http://localhost" + port)
	utils.OpenBrowser("http://localhost" + port)

	log.Fatal(http.ListenAndServe(port, mux))
}
