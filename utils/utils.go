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
		http.Error(w, "Error: invalid admin ID format", http.StatusBadRequest)
		return false
	}

	adminCheck, err := models.FindAdmin(context.Background(), db, adminIDInt)
	if err != nil {
		http.Error(w, "unable to verify admin ID", http.StatusServiceUnavailable)
		return false
	}

	userAdmin, err := models.FindUser(context.Background(), db, adminCheck.UserID)
	if err != nil {
		http.Error(w, "unable to find user associated with admin ID", http.StatusServiceUnavailable)
		return false
	}

	if userAdmin.Role != "admin" {
		http.Error(w, "user associated with provided admin ID is not an admin", http.StatusForbidden)
		return false
	}

	return true
}
