package main

import (
	"medical-clinic/api"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		return
	}
	api.StartServer()
}
