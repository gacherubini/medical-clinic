package utils

import (
	"context"
	"database/sql"
	"medical-clinic/models"
	"net/http"
	"strconv"
)

func IsAdminAllowed(w http.ResponseWriter, r *http.Request, db *sql.DB) bool {
	adminIDString := r.Header.Get("ID")

	adminIDInt, err := strconv.Atoi(adminIDString)
	if err != nil {
		http.Error(w, "Error in conversion", http.StatusServiceUnavailable)
		return false
	}

	adminCheck, err := models.FindAdmin(context.Background(), db, adminIDInt)
	if err != nil {
		http.Error(w, "Invalid ID, cant find admin", http.StatusServiceUnavailable)
		return false
	}

	userAdmin, err := models.FindUser(context.Background(), db, adminCheck.UserID)
	if err != nil {
		http.Error(w, "Invalid ID, cant find user from admin", http.StatusServiceUnavailable)
		return false
	}

	if userAdmin.Role != "admin" {
		http.Error(w, "Invalid ID, user isnt admin", http.StatusBadRequest)
		return false
	}

	return true
}
