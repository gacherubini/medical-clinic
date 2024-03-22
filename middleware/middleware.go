package middleware

import (
	"database/sql"
	"fmt"
	"medical-clinic/utils"
	"net/http"
)

type AdminMiddlewareContext struct {
	Db *sql.DB
}

func (AdminMiddlewareContext *AdminMiddlewareContext) IsAdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !utils.IsAdminAllowed(r, AdminMiddlewareContext.Db) {
			fmt.Fprintf(w, "Not an admin or is a diferent User")
			return
		}
		next.ServeHTTP(w, r)
	})
}
