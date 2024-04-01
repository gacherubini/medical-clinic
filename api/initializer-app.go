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

var subR = r.PathPrefix("").Subrouter()

const PORT = 8080

func StartServer() {
	adminMiddlewareContext := middleware.AdminMiddlewareContext{
		Db: Db,
	}

	subR.Use(adminMiddlewareContext.IsAdminMiddleware)

	for _, route := range getAdminRoutes() {
		if route.Method == http.MethodDelete || route.Method == http.MethodPatch {
			subR.HandleFunc(route.Path, route.Handler).Methods(route.Method)
		}
		r.HandleFunc(route.Path, route.Handler).Methods(route.Method)
	}

	for _, route := range getDoctorRoutes() {
		if route.Method == http.MethodDelete || route.Method == http.MethodPatch {
			subR.HandleFunc(route.Path, route.Handler).Methods(route.Method)
		}
		r.HandleFunc(route.Path, route.Handler).Methods(route.Method)
	}

	for _, route := range getPatientRoutes() {
		if route.Method == http.MethodDelete || route.Method == http.MethodPatch {
			subR.HandleFunc(route.Path, route.Handler).Methods(route.Method)
		}
		r.HandleFunc(route.Path, route.Handler).Methods(route.Method)
	}

	for _, route := range getAuthRoutes() {
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
