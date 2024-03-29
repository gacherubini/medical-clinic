package api

import (
	"database/sql"
	"fmt"
)

const (
	host     = "db"
	port     = 5432
	user     = "admin"
	password = "admin"
	dbname   = "postgres"
)

var psqlconn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", psqlconn)
	CheckError(err)
}
