package api

import (
	"context"
	"encoding/json"
	"fmt"
	"medical-clinic/models"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func HandleCreateDoctor(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
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
	user := userData.User
	doctor := userData.Doctor

	user.Insert(context.Background(), db, boil.Infer())
	doctor.Insert(context.Background(), db, boil.Infer())

	user.AddDoctors(context.Background(), db, true, &doctor)

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Doctor created successfully")
}

func HandleGetAllDoctors(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
		return
	}

	doctors, err := models.Doctors(qm.Load(models.DoctorRels.User)).All(context.Background(), db)
	if err != nil {
		http.Error(w, "Error retrieving doctors", http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(doctors)
	if err != nil {
		http.Error(w, "Error marshaling response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
