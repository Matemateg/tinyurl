package main

import (
	"database/sql"
	"github.com/Matemateg/tinyurl/database"
	"github.com/Matemateg/tinyurl/handlers"
	"github.com/Matemateg/tinyurl/handlers/api"
	"github.com/Matemateg/tinyurl/service"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
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
	srv := service.NewCreateUrl(base)
	http.HandleFunc("/", handlers.MainPage)
	http.Handle("/create", handlers.NewCreatingURL(srv))
	http.Handle("/t/", handlers.NewRedirectUrl(base))
	http.Handle("/api/create", api.NewCreatingURL(srv))
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
