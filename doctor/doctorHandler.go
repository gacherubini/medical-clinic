package doctor

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"medical-clinic/models"
	"medical-clinic/prepare"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"golang.org/x/crypto/bcrypt"
)

type userDoctorStruct struct {
	User   models.User
	Doctor models.Doctor
}

type DoctorHandlerContext struct {
	Db *sql.DB
}

func (contextHandler *DoctorHandlerContext) HandleCreateDoctor(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
		return
	}

	var userDoctor userDoctorStruct

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userDoctor); err != nil {
		http.Error(w, fmt.Sprintf("Error decoding JSON: %s", err), http.StatusBadRequest)
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(userDoctor.User.HashPassword), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error generating hashPassword: %s", err), http.StatusInternalServerError)
		return
	}

	userDoctor.User.HashPassword = string(passwordHash)

	if strings.ToLower(userDoctor.User.Role) != "doctor" {
		http.Error(w, "Invalid role, expected doctor", http.StatusBadRequest)
		return
	}

	err = userDoctor.User.Insert(context.Background(), contextHandler.Db, boil.Infer())
	if err != nil {
		http.Error(w, fmt.Sprintf("Error inserting user: %s", err), http.StatusInternalServerError)
		return
	}

	userDoctor.Doctor.UserID = userDoctor.User.UserID

	err = userDoctor.Doctor.Insert(context.Background(), contextHandler.Db, boil.Infer())
	if err != nil {
		http.Error(w, fmt.Sprintf("Error inserting doctor: %s", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Doctor created successfully")
}

func (contextHandler *DoctorHandlerContext) HandleGetAllDoctors(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
		return
	}

	doctors, err := models.Doctors(qm.Load(models.DoctorRels.User), qm.Load(models.DoctorRels.Healthinsurance)).All(context.Background(), contextHandler.Db)
	if err != nil {
		http.Error(w, "Error retrieving doctors", http.StatusInternalServerError)
		return
	}

	combinedData := prepare.PrepareDoctor(doctors)

	jsonDoctors, err := json.Marshal(combinedData)
	if err != nil {
		http.Error(w, "Error marshaling response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonDoctors)
}

func (contextHandler *DoctorHandlerContext) HandleDeleteDoctor(w http.ResponseWriter, r *http.Request) {
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

	doctor, err := models.FindDoctor(context.Background(), contextHandler.Db, intID)
	if err != nil {
		http.Error(w, "Failed to retrieve Doctor", http.StatusInternalServerError)
		return
	}

	doctorUser, err := models.FindUser(context.Background(), contextHandler.Db, doctor.UserID)
	if err != nil {
		http.Error(w, "Failed to retrieve User", http.StatusInternalServerError)
		return
	}

	if strings.ToLower(doctorUser.Role) != "doctor" {
		http.Error(w, "Invalid role, expected doctor", http.StatusBadRequest)
		return
	}

	if doctor.DoctorID != intID {
		http.Error(w, "Invalid ID, cant delete others doctors", http.StatusBadRequest)
		return
	}

	_, err = doctor.Delete(context.Background(), contextHandler.Db)
	if err != nil {
		http.Error(w, "Error deleting this Doctor", http.StatusInternalServerError)
	}

	_, err = doctorUser.Delete(context.Background(), contextHandler.Db)
	if err != nil {
		http.Error(w, "Error deleting this User from Doctor", http.StatusInternalServerError)
	}

	fmt.Fprintf(w, "Doctor deleted successfully")
}

func (contextHandler *DoctorHandlerContext) HandlerUpdateDoctor(w http.ResponseWriter, r *http.Request) {
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

	doctor, err := models.FindDoctor(context.Background(), contextHandler.Db, intID)
	if err != nil {
		http.Error(w, "Failed to retrieve doctor", http.StatusInternalServerError)
		return
	}

	doctorUser, err := models.FindUser(context.Background(), contextHandler.Db, doctor.UserID)
	if err != nil {
		http.Error(w, "Failed to retrieve User", http.StatusInternalServerError)
		return
	}

	if strings.ToLower(doctorUser.Role) != "doctor" {
		http.Error(w, "Invalid role, expected doctor", http.StatusBadRequest)
		return
	}

	doctorToUpdate := doctor

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&doctorToUpdate); err != nil {
		http.Error(w, "Error to Decode JSON", http.StatusBadRequest)
		return
	}

	_, err = doctorToUpdate.Update(context.Background(), contextHandler.Db, boil.Infer())
	if err != nil {
		http.Error(w, "Error updating doctor", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "doctor updated successfully")
}

func (contextHandler *DoctorHandlerContext) HandlerAddHealthInsurenceInDoctor(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
		return
	}

	params := mux.Vars(r)
	id := params["id"]
	intID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid doctor ID", http.StatusBadRequest)
		return
	}

	doctor, err := models.FindDoctor(context.Background(), contextHandler.Db, intID)
	if err != nil {
		http.Error(w, "Failed to retrieve Doctor", http.StatusInternalServerError)
		return
	}

	var healthinsurance models.Healthinsurance

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&healthinsurance); err != nil {
		http.Error(w, fmt.Sprintf("Error decoding JSON: %s", err), http.StatusBadRequest)
		return
	}

	if err := healthinsurance.Insert(context.Background(), contextHandler.Db, boil.Infer()); err != nil {
		http.Error(w, fmt.Sprintf("Error inserting new health insurance: %s", err), http.StatusInternalServerError)
		return
	}

	doctor.HealthinsuranceID = null.Int{Int: healthinsurance.HealthinsuranceID, Valid: true}
	doctor.Update(context.Background(), contextHandler.Db, boil.Infer())

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "HealthInsurence added in doctor successfully")
}
