package api

import (
	"context"
	"encoding/json"
	"fmt"
	"medical-clinic/models"
	"medical-clinic/prepare"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func HandleCreatePatient(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
		return
	}

	var userData struct {
		User    models.User    `json:"user"`
		Patient models.Patient `json:"patient"`
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userData); err != nil {
		http.Error(w, fmt.Sprintf("Error decoding JSON: %s", err), http.StatusBadRequest)
		return
	}

	if strings.ToLower(userData.User.Role) != "patient" {
		http.Error(w, "Invalid role, expected patient", http.StatusBadRequest)
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

	userData.Patient.UserID = userData.User.UserID

	err = userData.Patient.Insert(context.Background(), tx, boil.Infer())
	if err != nil {
		http.Error(w, fmt.Sprintf("Error inserting Patient: %s", err), http.StatusInternalServerError)
		return
	}
	err = tx.Commit()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error committing transaction: %s", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Patient created successfully")
}

func HandleGetAllPatient(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
		return
	}
	patients, err := models.Patients(qm.Load(models.PatientRels.User), qm.Load(models.PatientRels.Healthinsurance)).All(context.Background(), db)
	if err != nil {
		http.Error(w, "Error retrieving patients", http.StatusInternalServerError)
		return
	}

	combinedData := prepare.PreparePatient(patients)

	jsonPatient, err := json.Marshal(combinedData)
	if err != nil {
		http.Error(w, "Error marshaling response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonPatient)
}

func HandleDeletePatient(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
		return
	}

	params := mux.Vars(r)
	id := params["id"]
	intID, err := strconv.Atoi(id)

	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	patient, err := models.FindPatient(context.Background(), db, intID)
	if err != nil {
		http.Error(w, "Failed to retrieve patient", http.StatusInternalServerError)
		return
	}

	patientUser, err := models.FindUser(context.Background(), db, patient.UserID)
	if err != nil {
		http.Error(w, "Failed to retrieve User", http.StatusInternalServerError)
		return
	}

	if strings.ToLower(patientUser.Role) != "patient" {
		http.Error(w, "Invalid role, expected patient", http.StatusBadRequest)
		return
	}

	_, err = patient.Delete(context.Background(), db)
	if err != nil {
		http.Error(w, "Error deleting this patient", http.StatusInternalServerError)
	}

	_, err = patientUser.Delete(context.Background(), db)
	if err != nil {
		http.Error(w, "Error deleting this User from patient", http.StatusInternalServerError)
	}

	fmt.Fprintf(w, "Patient deleted successfully")
}

func HandlerUpdatePatient(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
		return
	}

	params := mux.Vars(r)
	id := params["id"]
	intID, err := strconv.Atoi(id)

	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	patient, err := models.FindPatient(context.Background(), db, intID)
	if err != nil {
		http.Error(w, "Failed to retrieve patient", http.StatusInternalServerError)
		return
	}

	patientUser, err := models.FindUser(context.Background(), db, patient.UserID)
	if err != nil {
		http.Error(w, "Failed to retrieve User", http.StatusInternalServerError)
		return
	}

	if strings.ToLower(patientUser.Role) != "patient" {
		http.Error(w, "Invalid role, expected patient", http.StatusBadRequest)
		return
	}

	patientToUpdate := patient

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&patientToUpdate); err != nil {
		http.Error(w, "Error to Decode JSON", http.StatusBadRequest)
		return
	}

	_, err = patientToUpdate.Update(context.Background(), db, boil.Infer())
	if err != nil {
		http.Error(w, "Error updating patient", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "patient updated successfully")
}

func HandlerAddHealthInsurenceInPatient(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
		return
	}

	params := mux.Vars(r)
	id := params["id"]
	intID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid patient ID", http.StatusBadRequest)
		return
	}

	patient, err := models.FindPatient(context.Background(), db, intID)
	if err != nil {
		http.Error(w, "Failed to retrieve Patient", http.StatusInternalServerError)
		return
	}

	var healthinsurance models.Healthinsurance

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&healthinsurance); err != nil {
		http.Error(w, fmt.Sprintf("Error decoding JSON: %s", err), http.StatusBadRequest)
		return
	}

	if err := healthinsurance.Insert(context.Background(), db, boil.Infer()); err != nil {
		http.Error(w, fmt.Sprintf("Error inserting new health insurance: %s", err), http.StatusInternalServerError)
		return
	}

	patient.HealthinsuranceID = null.Int{Int: healthinsurance.HealthinsuranceID, Valid: true}
	patient.Update(context.Background(), db, boil.Infer())

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "HealthInsurence added in patient successfully")
}
