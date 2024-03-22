package utils

import (
	"context"
	"database/sql"
	"medical-clinic/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func IsAdminAllowed(r *http.Request, db *sql.DB) bool {
	adminIDString := r.Header.Get("ID")

	if adminIDString == "" {
		return true
	}

	adminIDInt, err := strconv.Atoi(adminIDString)
	if err != nil {
		return false
	}

	if isUserAllowed(r, db) {
		return true
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

func isUserAllowed(r *http.Request, db *sql.DB) bool {
	params := mux.Vars(r)
	id := params["id"]
	intID, err := strconv.Atoi(id)

	if err != nil {
		return false
	}

	user, err := models.FindUser(context.Background(), db, intID)

	switch user.Role {
	case "doctor":
		doctor, _ := models.FindDoctor(context.Background(), db, user.UserID)
		if doctor.DoctorID == intID {
			return true
		}
	case "patient":
		patient, _ := models.FindPatient(context.Background(), db, user.UserID)
		if patient.PatientID == intID {
			return true
		}
	}
	return false
}
