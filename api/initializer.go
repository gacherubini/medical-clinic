package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

const (
	host     = "db"
	port     = 5432
	user     = "admin"
	password = "admin"
	dbname   = "postgres"
)

const PORT = 8080

var db *sql.DB

var psqlconn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

func init() {
	var err error
	db, err = sql.Open("postgres", psqlconn)
	CheckError(err)

	defer db.Close()
}

func StartServer() {
	for _, route := range routes {
		http.HandleFunc(route.Path, route.Handler)
	}

	fmt.Printf("Starting server at port %d\n", PORT)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil); err != nil {
		log.Fatal(err)
	}
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
