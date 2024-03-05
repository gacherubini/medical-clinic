package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

func main() {
	const (
		host     = "localhost"
		port     = 5432
		user     = "admin"
		password = "admin"
		dbname   = "postgres"
	)
	dbString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", dbString)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatal("specify a command (up, down, status, create)")
	}
	command, args := args[0], args[1:]
	if err := goose.RunContext(context.Background(), command, db, ".", args...); err != nil {
		log.Fatalf("goose run: %v", err)
	}
	fmt.Println("Goose command ran successfully!")
}
