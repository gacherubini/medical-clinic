package api

import (
	"fmt"
	"log"
	"net/http"

	"medical-clinic/middleware"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var r = mux.NewRouter()

const PORT = 8080

func StartServer() {
	adminMiddlewareContext := middleware.AdminMiddlewareContext{
		Db: Db,
	}

	r.Use(adminMiddlewareContext.IsAdminMiddleware)

	for _, route := range getAdminRoutes() {
		r.HandleFunc(route.Path, route.Handler).Methods(route.Method)
	}

	for _, route := range GetDoctorRoutes() {
		r.HandleFunc(route.Path, route.Handler).Methods(route.Method)
	}

	for _, route := range getPatientRoutes() {
		r.HandleFunc(route.Path, route.Handler).Methods(route.Method)
	}

	fmt.Printf("Starting server at port %d\n", PORT)

	defer Db.Close()

	if err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), r); err != nil {
		log.Fatal(err)
	}
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
