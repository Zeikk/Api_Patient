package main

import (
	"fmt"
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"./router"
)

func main() {
	fmt.Println("Hello, world.")

	db, err := sql.Open("mysql",
		"root:root@tcp(127.0.0.1:3306)/api_go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nConnexion à la base")

	router := router.Router()

	fmt.Println("\nDémarrage du serveur port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
	defer db.Close()
}