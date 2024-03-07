package api

import (
	"context"
	"encoding/json"
	"fmt"
	"medical-clinic/models"
	"net/http"
	"strings"

	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func HandleCreateDoctor(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
		return
	}

	var userData struct {
		User   models.User   `json:"user"`
		Doctor models.Doctor `json:"doctor"`
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userData); err != nil {
		http.Error(w, fmt.Sprintf("Error decoding JSON: %s", err), http.StatusBadRequest)
		return
	}

	if strings.ToLower(userData.User.Role) != "doctor" {
		http.Error(w, "Invalid role, expected doctor", http.StatusBadRequest)
		return
	}

	tx, err := db.Begin()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error starting transaction: %s", err), http.StatusInternalServerError)
		return
	}
	defer tx.Rollback()

	err = userData.User.Insert(context.Background(), tx, boil.Infer())
	if err != nil {
		http.Error(w, fmt.Sprintf("Error inserting user: %s", err), http.StatusInternalServerError)
		return
	}

	userData.Doctor.UserID = userData.User.UserID

	err = userData.Doctor.Insert(context.Background(), tx, boil.Infer())
	if err != nil {
		http.Error(w, fmt.Sprintf("Error inserting doctor: %s", err), http.StatusInternalServerError)
		return
	}

	err = tx.Commit()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error committing transaction: %s", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Doctor created successfully")
}

func HandleGetAllDoctors(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
		return
	}

	doctors, err := models.Doctors(qm.Load(models.DoctorRels.User)).All(context.Background(), db)
	if err != nil {
		http.Error(w, "Error retrieving doctors", http.StatusInternalServerError)
		return
	}

	var combinedData []map[string]interface{}

	for _, doctor := range doctors {
		user := doctor.R.User

		responseData := map[string]interface{}{
			"doctor": doctor,
			"user":   user,
		}

		combinedData = append(combinedData, responseData)
	}

	jsonDoctors, err := json.Marshal(combinedData)
	if err != nil {
		http.Error(w, "Error marshaling response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonDoctors)
}
