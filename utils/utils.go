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
		return false
	}

	adminCheck, err := models.FindAdmin(context.Background(), db, adminIDInt)
	if err != nil {
		return false
	}

	userAdmin, err := models.FindUser(context.Background(), db, adminCheck.UserID)
	if err != nil {
		return false
	}

	if userAdmin.Role != "admin" {
		return false
	}

	return true
}
