package api

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

const PORT = 8080

func StartServer() {
	for _, route := range routes {
		http.HandleFunc(route.Path, route.Handler)
	}

	fmt.Printf("Starting server at port %d\n", PORT)
	defer db.Close()
	if err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil); err != nil {
		log.Fatal(err)
	}
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
