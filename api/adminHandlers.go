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
	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func HandleCreateAdmin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
		return
	}

	var userData struct {
		User  models.User  `json:"user"`
		Admin models.Admin `json:"admin"`
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userData); err != nil {
		http.Error(w, fmt.Sprintf("Error decoding JSON: %s", err), http.StatusBadRequest)
		return
	}

	if strings.ToLower(userData.User.Role) != "admin" {
		http.Error(w, "Invalid role, expected admin", http.StatusBadRequest)
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

	userData.Admin.UserID = userData.User.UserID

	err = userData.Admin.Insert(context.Background(), tx, boil.Infer())
	if err != nil {
		http.Error(w, fmt.Sprintf("Error inserting admin: %s", err), http.StatusInternalServerError)
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

func HandleGetAllAdmins(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
		return
	}

	admins, err := models.Admins(qm.Load(models.DoctorRels.User)).All(context.Background(), db)
	if err != nil {
		http.Error(w, "Error retrieving admins", http.StatusInternalServerError)
		return
	}

	combinedData := prepare.PrepareAdmin(admins)

	jsonDoctors, err := json.Marshal(combinedData)
	if err != nil {
		http.Error(w, "Error marshaling response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonDoctors)
}

func HandleDeleteAdmin(w http.ResponseWriter, r *http.Request) {
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

	admin, err := models.FindAdmin(context.Background(), db, intID)
	if err != nil {
		http.Error(w, "Failed to retrieve Admin", http.StatusInternalServerError)
		return
	}

	AdminUser, err := models.FindUser(context.Background(), db, admin.UserID)
	if err != nil {
		http.Error(w, "Failed to retrieve User", http.StatusInternalServerError)
		return
	}

	if strings.ToLower(AdminUser.Role) != "admin" {
		http.Error(w, "Invalid role, expected admin", http.StatusBadRequest)
		return
	}

	_, err = admin.Delete(context.Background(), db)
	if err != nil {
		http.Error(w, "Error deleting this admin", http.StatusInternalServerError)
	}

	_, err = AdminUser.Delete(context.Background(), db)
	if err != nil {
		http.Error(w, "Error deleting this User from Admin", http.StatusInternalServerError)
	}

	fmt.Fprintf(w, "Admin deleted successfully")
}

func HandleUpdateAdmin(w http.ResponseWriter, r *http.Request) {
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

	admin, err := models.FindAdmin(context.Background(), db, intID)
	if err != nil {
		http.Error(w, "Failed to retrieve admin", http.StatusInternalServerError)
		return
	}

	adminUser, err := models.FindUser(context.Background(), db, admin.UserID)
	if err != nil {
		http.Error(w, "Failed to retrieve User", http.StatusInternalServerError)
		return
	}

	if strings.ToLower(adminUser.Role) != "admin" {
		http.Error(w, "Invalid role, expected admin", http.StatusBadRequest)
		return
	}

	adminToUpdate := admin

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&adminToUpdate); err != nil {
		http.Error(w, "Error to Decode JSON", http.StatusBadRequest)
		return
	}

	_, err = adminToUpdate.Update(context.Background(), db, boil.Infer())
	if err != nil {
		http.Error(w, "Error updating admin", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "admin updated successfully")
}
