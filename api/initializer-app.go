package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var r = mux.NewRouter()

const PORT = 8080

func StartServer() {
	for _, route := range routes {
		r.HandleFunc(route.Path, route.Handler)
	}

	fmt.Printf("Starting server at port %d\n", PORT)

	defer db.Close()

	if err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), r); err != nil {
		log.Fatal(err)
	}
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
