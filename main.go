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

	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/add", handlers.AddHandler)
	http.HandleFunc("/delete", handlers.DeleteHandler)
	http.HandleFunc("/update", handlers.UpdateHandler)
	http.HandleFunc("/edit", handlers.EditHandler)

	port := ":7777"
	fmt.Println("Servidor rodando em http://localhost" + port)
	utils.OpenBrowser("http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
