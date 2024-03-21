package middleware

import (
	"database/sql"
	"fmt"
	"medical-clinic/utils"
	"net/http"
)

func IsAdminMiddleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !utils.IsAdminAllowed(r, db) {
			fmt.Fprintf(w, "Not an admin or is a diferent User")
			return
		}
		next.ServeHTTP(w, r)
	})
}
