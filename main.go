package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"github.com/Matemateg/tinyurl/database"
	"github.com/Matemateg/tinyurl/handlers"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8090"
	}

	mysqlDSN := os.Getenv("MYSQL_DSN")
	if mysqlDSN == "" {
		mysqlDSN = "root:123@/tinyurls"
	}

	db, err := sql.Open("mysql", mysqlDSN)
	if err != nil {
		log.Fatalln(err)
	}

	var base = database.New(db)
	http.HandleFunc("/", handlers.MainPage)
	http.Handle("/create", handlers.NewCreateUrl(base))
	http.Handle("/t/", handlers.NewRedirectUrl(base))
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
