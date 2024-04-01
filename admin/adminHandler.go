package admin

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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"golang.org/x/crypto/bcrypt"
)

type UserAdminStruct struct {
	User  models.User
	Admin models.Admin
}

type AdminHandlerContext struct {
	Db *sql.DB
}

func (contextHandler *AdminHandlerContext) HandleCreateAdmin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
		return
	}

	var userAdmin UserAdminStruct

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userAdmin); err != nil {
		http.Error(w, fmt.Sprintf("Error decoding JSON: %s", err), http.StatusBadRequest)
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(userAdmin.User.HashPassword), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error generating hashPassword: %s", err), http.StatusInternalServerError)
		return
	}

	userAdmin.User.HashPassword = string(passwordHash)

	if strings.ToLower(userAdmin.User.Role) != "admin" {
		http.Error(w, "Invalid role, expected admin", http.StatusBadRequest)
		return
	}

	err = userAdmin.User.Insert(context.Background(), contextHandler.Db, boil.Infer())
	if err != nil {
		http.Error(w, fmt.Sprintf("Error inserting user: %s", err), http.StatusInternalServerError)
		return
	}

	userAdmin.Admin.UserID = userAdmin.User.UserID

	err = userAdmin.Admin.Insert(context.Background(), contextHandler.Db, boil.Infer())
	if err != nil {
		http.Error(w, fmt.Sprintf("Error inserting admin: %s", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Admin created successfully")
}

func (contextHandler *AdminHandlerContext) HandleGetAllAdmins(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
		return
	}

	admins, err := models.Admins(qm.Load(models.AdminRels.User)).All(context.Background(), contextHandler.Db)
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

func (contextHandler *AdminHandlerContext) HandleDeleteAdmin(w http.ResponseWriter, r *http.Request) {
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

	admin, err := models.FindAdmin(context.Background(), contextHandler.Db, intID)
	if err != nil {
		http.Error(w, "Failed to retrieve Admin", http.StatusInternalServerError)
		return
	}

	AdminUser, err := models.FindUser(context.Background(), contextHandler.Db, admin.UserID)
	if err != nil {
		http.Error(w, "Failed to retrieve User", http.StatusInternalServerError)
		return
	}

	if strings.ToLower(AdminUser.Role) != "admin" {
		http.Error(w, "Invalid role, expected admin", http.StatusBadRequest)
		return
	}

	_, err = admin.Delete(context.Background(), contextHandler.Db)
	if err != nil {
		http.Error(w, "Error deleting this admin", http.StatusInternalServerError)
	}

	_, err = AdminUser.Delete(context.Background(), contextHandler.Db)
	if err != nil {
		http.Error(w, "Error deleting this User from Admin", http.StatusInternalServerError)
	}

	fmt.Fprintf(w, "Admin deleted successfully")
}

func (contextHandler *AdminHandlerContext) HandleUpdateAdmin(w http.ResponseWriter, r *http.Request) {
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

	admin, err := models.FindAdmin(context.Background(), contextHandler.Db, intID)
	if err != nil {
		http.Error(w, "Failed to retrieve admin", http.StatusInternalServerError)
		return
	}

	adminUser, err := models.FindUser(context.Background(), contextHandler.Db, admin.UserID)
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

	_, err = adminToUpdate.Update(context.Background(), contextHandler.Db, boil.Infer())
	if err != nil {
		http.Error(w, "Error updating admin", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "admin updated successfully")
}

func (contextHandler *AdminHandlerContext) HandleAdminCreateHealthInsurence(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
		return
	}

	var healthinsurance models.Healthinsurance

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&healthinsurance); err != nil {
		http.Error(w, fmt.Sprintf("Error decoding JSON: %s", err), http.StatusBadRequest)
		return
	}

	err := healthinsurance.Insert(context.Background(), contextHandler.Db, boil.Infer())
	if err != nil {
		http.Error(w, fmt.Sprintf("Error inserting user: %s", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "HealthInsurance created successfully")
}

func (contextHandler *AdminHandlerContext) HandleAdminGetAllHealthInsurence(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
		return
	}

	healthInsurance, err := models.Healthinsurances().All(context.Background(), contextHandler.Db)

	jsonDoctors, err := json.Marshal(healthInsurance)
	if err != nil {
		http.Error(w, "Error marshaling response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonDoctors)
}
